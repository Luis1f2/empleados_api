package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Luis1f2/empleados_api/src/empleados/application"
)

type EmpleadoController struct {
	createEmpleado  *application.CreateEmpleado
	getEmpleadoByID *application.GetEmpleadoByID
	listEmpleados   *application.ListEmpleados
	updateEmpleado  *application.UpdateEmpleado
	deleteEmpleado  *application.DeleteEmpleado
}

func NewEmpleadoController(
	createEmpleado *application.CreateEmpleado,
	getEmpleadoByID *application.GetEmpleadoByID,
	listEmpleados *application.ListEmpleados,
	updateEmpleado *application.UpdateEmpleado,
	deleteEmpleado *application.DeleteEmpleado,
) *EmpleadoController {
	return &EmpleadoController{
		createEmpleado:  createEmpleado,
		getEmpleadoByID: getEmpleadoByID,
		listEmpleados:   listEmpleados,
		updateEmpleado:  updateEmpleado,
		deleteEmpleado:  deleteEmpleado,
	}
}

func (ec *EmpleadoController) Create(c *gin.Context) {
	var body struct {
		Nombre    string `json:"nombre"`
		Apellidos string `json:"apellidos"`
		Mail      string `json:"mail"`
		User      string `json:"username"`
		Password  string `json:"password"`
		RolID     int    `json:"rol_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"})
		return
	}

	empleado, err := ec.createEmpleado.Execute(
		body.Nombre,
		body.Apellidos,
		body.Mail,
		body.User,
		body.Password,
		body.RolID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, empleado)
}

func (ec *EmpleadoController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	empleado, err := ec.getEmpleadoByID.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "empleado no encontrado"})
		return
	}

	c.JSON(http.StatusOK, empleado)
}

func (ec *EmpleadoController) List(c *gin.Context) {
	empleados, err := ec.listEmpleados.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error al listar empleados"})
		return
	}

	c.JSON(http.StatusOK, empleados)
}

func (ec *EmpleadoController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	var body struct {
		Nombre    string `json:"nombre"`
		Apellidos string `json:"apellidos"`
		Mail      string `json:"mail"`
		User      string `json:"username"`
		Password  string `json:"password"`
		RolID     int    `json:"rol_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"})
		return
	}

	empleado, err := ec.updateEmpleado.Execute(
		id,
		body.Nombre,
		body.Apellidos,
		body.Mail,
		body.User,
		body.Password,
		body.RolID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, empleado)
}

func (ec *EmpleadoController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	err = ec.deleteEmpleado.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "empleado eliminado correctamente",
	})
}