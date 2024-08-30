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
	// Return the new customer in the response
	c.JSON(http.StatusOK, gin.H{
		"message":  "Item added successfully",
		"customer": customer,
	})

}

func CustomerAuthentication(c *gin.Context) {
	var body struct {
		Customer_email string `json:"customer_email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Retrieve the customer
	var customer []models.Customer
	if err := initializers.DB.Where("customer_email = ?", body.Customer_email).Find(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "customer not found",
		})
		return
	}
	// Return the customer
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"cart":    customer,
	})
}
