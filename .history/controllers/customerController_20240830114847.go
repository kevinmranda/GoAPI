package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	// Create a new cart for the customer
	customer := models.Customer{
		CustomerEmail: body.Customer_email,
	}

}

func CustomerAuthentication(c *gin.Context) {

}
