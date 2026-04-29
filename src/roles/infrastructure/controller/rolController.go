package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Luis1f2/empleados_api/src/roles/application"
)

type RolController struct {
	createRole  *application.CreateRole
	getRoleByID *application.GetRoleByID
	listRoles   *application.ListRoles
	updateRole  *application.UpdateRole
	deleteRole  *application.DeleteRole
}

func NewRolController(
	createRole *application.CreateRole,
	getRoleByID *application.GetRoleByID,
	listRoles *application.ListRoles,
	updateRole *application.UpdateRole,
	deleteRole *application.DeleteRole,
) *RolController {
	return &RolController{
		createRole:  createRole,
		getRoleByID: getRoleByID,
		listRoles:   listRoles,
		updateRole:  updateRole,
		deleteRole:  deleteRole,
	}
}

func (rc *RolController) Create(c *gin.Context) {
	var body struct {
		Nombre string `json:"nombre"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"})
		return
	}

	role, err := rc.createRole.Execute(body.Nombre)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func (rc *RolController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	role, err := rc.getRoleByID.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "rol no encontrado"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (rc *RolController) List(c *gin.Context) {
	roles, err := rc.listRoles.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error al listar roles"})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func (rc *RolController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	var body struct {
		Nombre string `json:"nombre"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"})
		return
	}

	role, err := rc.updateRole.Execute(id, body.Nombre)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (rc *RolController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	err = rc.deleteRole.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "rol eliminado correctamente",
	})
}