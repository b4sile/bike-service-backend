package controllers

import (
	"net/http"
	"strconv"

	"github.com/b4sile/bike-service-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(len(users)))
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var user models.User
	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Select("Requests", "Notifications").Delete(&user)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value := map[string]interface{}{
		"Name":     input.Name,
		"Lastname": input.Lastname,
		"Surname":  input.Surname,
		"Email":    input.Email,
		"Phone":    input.Phone,
		"IsAdmin":  input.IsAdmin}

	models.DB.Model(&user).Updates(value)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     input.Name,
		Lastname: input.Lastname,
		Surname:  input.Surname,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		IsAdmin:  input.IsAdmin}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}
