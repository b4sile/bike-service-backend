package controllers

import (
	"net/http"
	"strconv"

	"github.com/b4sile/bike-service-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) {
	var services []models.Service
	if err := models.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(len(services)))
	c.JSON(http.StatusOK, services)
}

func GetService(c *gin.Context) {
	var service models.Service
	if err := models.DB.First(&service, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service)
}

func DeleteService(c *gin.Context) {
	var service models.Service
	if err := models.DB.First(&service, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&service)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func UpdateService(c *gin.Context) {
	var service models.Service
	if err := models.DB.First(&service, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.Service
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value := map[string]interface{}{
		"Name":        input.Name,
		"Price":       input.Price,
		"Description": input.Description}

	models.DB.Model(&service).Updates(value)
	c.JSON(http.StatusOK, service)
}

func CreateService(c *gin.Context) {
	var input models.Service
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := models.Service{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description}
	models.DB.Create(&service)

	c.JSON(http.StatusOK, service)
}
