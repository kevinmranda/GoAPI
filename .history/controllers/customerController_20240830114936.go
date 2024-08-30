package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func AddCustomer(c *gin.Context) {
	var body struct {
		Customer_email string `json:"customer_email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Create a new customer
	customer := models.Customer{
		CustomerEmail: body.Customer_email,
	}
	if err := initializers.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    2000,
			"error": "Failed to insert the record",
		})
		return
	}

}

func CustomerAuthentication(c *gin.Context) {

}
