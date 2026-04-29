package middleware

import (
	"net/http"
	"strings"

	authInfra "github.com/Luis1f2/empleados_api/src/auth/infrastructure"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token requerido"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := authInfra.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		c.Set("empleado_id", claims.EmpleadoID)
		c.Set("user", claims.User)
		c.Set("rol_id", claims.RolID)

		c.Next()
	}
}

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		rolIDValue, exists := c.Get("rol_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "rol no encontrado en token"})
			c.Abort()
			return
		}

		rolID, ok := rolIDValue.(int)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "rol inválido"})
			c.Abort()
			return
		}

		if rolID != 1 {
			c.JSON(http.StatusForbidden, gin.H{"error": "solo admin puede realizar esta acción"})
			c.Abort()
			return
		}

		c.Next()
	}
}