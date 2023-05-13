package middlewares

import (
	"fmt"
	"net/http"

	"github.com/b4sile/bike-service-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("admin AdminAuthMiddleware")
		var user models.User
		if err := models.DB.First(&user, c.GetUint("userId")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !user.IsAdmin {
			c.String(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}
		c.Next()
	}
}
