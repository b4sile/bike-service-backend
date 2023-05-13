package controllers

import (
	"net/http"

	"github.com/b4sile/bike-service-backend/internal/helpers"
	"github.com/b4sile/bike-service-backend/internal/models"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "msg": "not valid input"})
		return
	}

	var user models.User
	if err := models.DB.First(&user, "email = ?", input.Email).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordIsValid := helpers.CheckPassword(user.Password, input.Password)

	if !passwordIsValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	token, err := helpers.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})

}
