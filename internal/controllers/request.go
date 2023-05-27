package controllers

import (
	"net/http"
	"strconv"

	"github.com/b4sile/bike-service-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func GetRequests(c *gin.Context) {
	var requests []models.Request
	mp := map[string]interface{}{}
	userId := c.Query("userId")
	if userId != "" {
		mp["user_id"] = userId
	}
	if err := models.DB.Where(mp).Find(&requests).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(len(requests)))
	c.JSON(http.StatusOK, requests)
}

func GetRequest(c *gin.Context) {
	var request models.Request
	if err := models.DB.First(&request, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

func DeleteRequest(c *gin.Context) {
	var request models.Request
	if err := models.DB.First(&request, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&request)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func UpdateRequest(c *gin.Context) {
	var request models.Request
	if err := models.DB.First(&request, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value := map[string]interface{}{
		"UserId":         input.UserId,
		"Status":         input.Status,
		"Latitude":       input.Latitude,
		"Longitude":      input.Longitude,
		"Phone":          input.Phone,
		"IsOuterRequest": input.IsOuterRequest,
		"Description":    input.Description}

	models.DB.Model(&request).Updates(value)
	c.JSON(http.StatusOK, request)
}

func CreateRequest(c *gin.Context) {
	var input models.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request := models.Request{
		UserId:         input.UserId,
		Status:         input.Status,
		Phone:          input.Phone,
		Latitude:       input.Latitude,
		Longitude:      input.Longitude,
		IsOuterRequest: input.IsOuterRequest,
		Description:    input.Description}
	models.DB.Create(&request)

	c.JSON(http.StatusOK, request)
}
