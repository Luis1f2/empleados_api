package routes

import (
	authMiddleware "github.com/Luis1f2/empleados_api/src/auth/infrastructure/middleware"
	"github.com/Luis1f2/empleados_api/src/roles/infrastructure/controller"
	"github.com/gin-gonic/gin"
)

func RolRoutes(router *gin.Engine, rolController *controller.RolController) {
	roles := router.Group("/api/roles")
	roles.Use(authMiddleware.AuthMiddleware())
	roles.Use(authMiddleware.RequireAdmin())

	{
		roles.POST("/", rolController.Create)
		roles.GET("/", rolController.List)
		roles.GET("/:id", rolController.GetByID)
		roles.PUT("/:id", rolController.Update)
		roles.DELETE("/:id", rolController.Delete)
	}
}