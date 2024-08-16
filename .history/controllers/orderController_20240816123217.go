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

	// if ValidateEmail(body.Customer_email) {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"id":    2002,
	// 		"error": "Invalid email format",
	// 	})
	// } else {
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Retrieve the associated photos based on PhotoIDs
	var photos []models.Photo
	if err := initializers.DB.Where("id IN ?", body.PhotoIDs).Find(&photos).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "Failed to retrieve photos",
		})
		return
	}

	// Create the order record
	order := models.Order{
		Customer_email: body.Customer_email,
		Total_amount:   body.Total_amount,
		Status:         body.Status,
		Photos:         photos,
	}

	if err := initializers.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    2000,
			"error": "Failed to insert the record",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "Order created successfully",
		"order":   order,
	})
	// }

}

// delete order
func RemoveOrder(c *gin.Context) {
	// Get id from request
	id := c.Param("id")

	var order models.Order
	// Check if the photo exists
	initializers.DB.First(&order, id)

	if order.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not found",
		})
		return
	}

	// Delete the order
	result := initializers.DB.Delete(&order)

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

// retrieve an order with id
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var orders []models.Order
	// Preload the many-to-many relationship with Photos and the one-to-one relationship with Payment
	result := initializers.DB.Preload("Photos").Preload("Payment").First(&orders, id)
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
		"data":    orders,
	})
}

// retrieve all orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	// Preload the many-to-many relationship with Photos and the one-to-one relationship with Payment
	result := initializers.DB.Preload("Photos").Preload("Payment").Find(&orders)
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
		"data":    orders,
	})
}

// update an order with id
func UpdateOrder(c *gin.Context) {
	// Get id from request
	id := c.Param("id")

	//struct for contents to be updated
	var contentForUpdate struct {
		Customer_email string  `json:"customer_email" binding:"required"`
		Total_amount   float64 `json:"total_amount" binding:"required"`
		Status         string  `json:"status" binding:"required"`
		PhotoIDs       []uint  `json:"photo_ids" binding:"required"`
	}

	//Get contents from body of request and Bind JSON input to the struct
	if err := c.ShouldBindJSON(&contentForUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Check if the order to be updated exists
	var order models.Order
	result := initializers.DB.First(&order, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not found",
		})
		return
	}

	//update order with struct provided on request body
	if contentForUpdate.Customer_email != "" {
		order.Customer_email = contentForUpdate.Customer_email
	}

	if contentForUpdate.Total_amount != 0 {
		order.Total_amount = contentForUpdate.Total_amount
	}

	if contentForUpdate.Status != "" {
		order.Status = contentForUpdate.Status
	}

	
	if contentForUpdate.PhotoIDs != nil {
		order.Photos = photos
	}

	result = initializers.DB.Save(&order)
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
		"data":    order,
	})
}
