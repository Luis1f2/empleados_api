package routes

import (
	"github.com/Luis1f2/empleados_api/src/auth/infrastructure/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authController *controller.AuthController) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", authController.Login)
	}
}