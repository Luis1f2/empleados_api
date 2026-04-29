package controller

import (
	

	authInfra "github.com/Luis1f2/empleados_api/src/auth/infrastructure"
	empleadosRepository "github.com/Luis1f2/empleados_api/src/empleados/domain/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	empleadoRepo empleadosRepository.EmpleadoRepository
}

func NewAuthController(empleadoRepo empleadosRepository.EmpleadoRepository) *AuthController {
	return &AuthController{
		empleadoRepo: empleadoRepo,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var body struct {
		User     string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "body inválido"})
		return
	}

	empleado, err := ac.empleadoRepo.GetByUser(body.User)
	if err != nil {
		c.JSON(401, gin.H{"error": "credenciales inválidas"})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(empleado.Password),
		[]byte(body.Password),
	)

	if err != nil {
		c.JSON(401, gin.H{"error": "credenciales inválidas"})
		return
	}

	token, err := authInfra.GenerateToken(
		empleado.ID,
		empleado.User,
		empleado.RolID,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "error generando token"})
		return
	}

	c.JSON(200, gin.H{
		"message": "login exitoso",
		"token":   token,
	})
}