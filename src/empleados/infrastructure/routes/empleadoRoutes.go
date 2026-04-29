package routes

import (
	authMiddleware "github.com/Luis1f2/empleados_api/src/auth/infrastructure/middleware"
	"github.com/Luis1f2/empleados_api/src/empleados/infrastructure/controller"
	"github.com/gin-gonic/gin"
)

func EmpleadoRoutes(router *gin.Engine, empleadoController *controller.EmpleadoController) {
	empleados := router.Group("/api/empleados")
	empleados.Use(authMiddleware.AuthMiddleware())
	empleados.Use(authMiddleware.RequireAdmin())

	{
		empleados.POST("/", empleadoController.Create)
		empleados.GET("/", empleadoController.List)
		empleados.GET("/:id", empleadoController.GetByID)
		empleados.PUT("/:id", empleadoController.Update)
		empleados.DELETE("/:id", empleadoController.Delete)
	}
}