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

	amountToBePaid := order.Total_amount
	if body.Amount == amountToBePaid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Enter the correct amount please",
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

	err := initializers.DB.Create(&payment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create payment",
		})
		return
	} else {
		order.Status = "completed"
		initializers.DB.Save(&order)
		// Respond with success
		c.JSON(http.StatusOK, gin.H{
			"message": "Payment created successfully",
			"payment": payment,
		})
	}

}

func DeletePayment(c *gin.Context) {
	c.Get("user")
	id := c.Param("id")
	var payment models.Payment
	initializers.DB.Find(&payment, id)
	if payment.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not present",
		})
		return
	}

	// Delete the payment
	result := initializers.DB.Delete(&payment)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"id":    2013,
			"error": "failed to delete record",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2012,
		"message": "record deleted successfully",
	})
}

func UpdatePayment(c *gin.Context) {
	c.Get("user")
	// Get id from request
	id := c.Param("id")

	//struct for contents to be updated
	var contentForUpdate struct {
		OrderID        uint    `json:"order_id"`
		Amount         float64 `json:"amount"`
		Status         string  `json:"status"`         // "pending", "completed", "failed"
		Payment_method string  `json:"payment_method"` // "credit_card", "paypal"
		Transaction_id string  `json:"transaction_id"` // Identifier from payment gateway
	}

	//Get contents from body of request and Bind JSON input to the struct
	if err := c.ShouldBindJSON(&contentForUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Check if the payment to be updated exists
	var payment models.Payment
	result := initializers.DB.First(&payment, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not found",
		})
		return
	}

	//update payment with struct provided on request body
	if contentForUpdate.OrderID != 0 {
		payment.Order_id = contentForUpdate.OrderID
	}

	if contentForUpdate.Amount != 0 {
		payment.Amount = contentForUpdate.Amount
	}

	if contentForUpdate.Status != "" {
		payment.Status = contentForUpdate.Status
	}

	if contentForUpdate.Payment_method != "" {
		payment.Payment_method = contentForUpdate.Payment_method
	}

	if contentForUpdate.Transaction_id != "" {
		payment.Transaction_id = contentForUpdate.Transaction_id
	}

	result = initializers.DB.Save(&payment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    2014,
			"error": "Failed to update the record",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "success",
		"data":    payment,
	})
}

func GetPayment(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment
	result := initializers.DB.Find(&payment, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not present",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "success",
		"data":    payment,
	})

}

func GetPayments(c *gin.Context) {
	userID := c.Param("id")
	var payments []models.Payment
	result := initializers.DB.
		Joins("JOIN orders ON payments.order_id = orders.id").
		Joins("JOIN order_photos ON orders.id = order_photos.order_id").
		Joins("JOIN photos ON order_photos.photo_id = photos.id").
		Where("photos.user_id = ?", userID).
		Find(&payments)

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
		"data":    payments,
	})
}
