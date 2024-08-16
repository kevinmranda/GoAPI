package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

// add order to db
func AddOrder(c *gin.Context) {
	// Bind JSON input to the struct
	var body struct {
		Customer_email string  `json:"customer_email" binding:"required"`
		Total_amount   float64 `json:"total_amount" binding:"required"`
		Status         string  `json:"status" binding:"required"`
		PhotoIDs       []uint  `json:"photo_ids" binding:"required"`
	}

	if !ValidateEmail(body.Customer_email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    2002,
			"error": "Invalid email format",
		})
		return
	} else {

	}

	
}

// delete order
func RemoveOrder(c *gin.Context) {

}

// retrieve an order with id
func GetOrder(c *gin.Context) {

}

// retrieve all orders
func GetOrders(c *gin.Context) {

}

// update an order with id
func UpdateOrder(c *gin.Context) {

}
