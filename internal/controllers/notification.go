package controllers

import (
	"net/http"
	"strconv"

	"github.com/b4sile/bike-service-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func GetNotifications(c *gin.Context) {
	var notifications []models.Notification
	mp := map[string]interface{}{}
	userId := c.Query("userId")
	if userId != "" {
		mp["user_id"] = userId
	}
	if err := models.DB.Where(mp).Find(&notifications).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(len(notifications)))
	c.JSON(http.StatusOK, notifications)
}

func GetNotification(c *gin.Context) {
	var notification models.Notification
	if err := models.DB.First(&notification, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notification)
}

func DeleteNotification(c *gin.Context) {
	var notification models.Notification
	if err := models.DB.First(&notification, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&notification)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func UpdateNotification(c *gin.Context) {
	var notification models.Notification
	if err := models.DB.First(&notification, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.Notification
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value := map[string]interface{}{
		"UserId":      input.UserId,
		"IsReaded":    input.IsReaded,
		"Description": input.Description}

	models.DB.Model(&notification).Updates(value)
	c.JSON(http.StatusOK, notification)
}

func CreateNotification(c *gin.Context) {
	var input models.Notification
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification := models.Notification{
		UserId:      input.UserId,
		IsReaded:    input.IsReaded,
		Description: input.Description}
	models.DB.Create(&notification)

	c.JSON(http.StatusOK, notification)
}
