package middlewares

import (
	"net/http"

	"github.com/b4sile/bike-service-backend/internal/helpers"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.IsTokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		userId, _ := helpers.ExtractUserIdFromToken(c)
		c.Set("userId", userId)
		c.Next()
	}
}
