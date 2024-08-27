package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/models"
)

func getLogs(c *gin.Context) {
	var logs []models.ActivityLog
	result := initializers.DB.Find(&logs)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "records not present",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "success",
		"data":    logs,
	})
}