package controllers

import (
	"fmt"

	"github.com/b4sile/bike-service-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	var user models.User
	err2 := models.DB.Preload("Requests").First(&user, 100)
	fmt.Println(err2.Error, "err2")

	c.JSON(200, gin.H{
		"message": user,
	})
}
