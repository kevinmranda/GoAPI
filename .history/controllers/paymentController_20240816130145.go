package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func AddPayment(c *gin.Context) {
	// Bind JSON input to the struct
	var body struct {
		OrderID        uint    `json:"order_id" binding:"required"`
		Amount         float64 `json:"amount" binding:"required"`
		Status         string  `json:"status" binding:"required"`         // "pending", "completed", "failed"
		Payment_method string  `json:"payment_method" binding:"required"` // "credit_card", "paypal"
		Transaction_id string  `json:"transaction_id" binding:"required"` // Identifier from payment gateway
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Check if the order exists
	var order models.Order
	if err := initializers.DB.First(&order, body.OrderID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not found",
		})
		return
	}

	// Create the payment record and associate it with the order
	payment := models.Payment{
		Order_id:       body.OrderID,
		Amount:         body.Amount,
		Status:         body.Status,
		Payment_method: body.Payment_method,
		Transaction_id: body.Transaction_id,
	}

	if err := initializers.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create payment",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment created successfully",
		"payment": payment,
	})
}

func DeletePayment() {

}

func UpdatePayment() {

}

func Get