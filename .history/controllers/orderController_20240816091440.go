package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

// Add order in db
func AddOrder(c *gin.Context) {
	//struct of the request body
	var body struct {
		Customer_email string  `json:"title" binding:"required"` //(Email for the order, used for sending download links)
		Total_amount   float64 `json:"title" binding:"required"`
		Status         string  `json:"title" binding:"required"` //(enum: "pending", "completed", "canceled")
		// Photos         []Photo `json:"title" binding:"required"`//`gorm:"many2many:order_photos;"`
		// Payment        Payment `json:"title" binding:"required"`// One-to-One relationship with Payment
	}

	//Get contents from body of request and Bind JSON input to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//add photo to db
	photo := models.Order{
		Title:       body.Title,
		Description: body.Description,
		Filename:    body.Filename,
		Price:       body.Price,
		User_id:     body.User_id,
	}
	result := initializers.DB.Create(&photo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    2000,
			"error": "Failed to insert the record",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"id":      2001,
		"message": "record inserted successfully",
		"data":    photo,
	})
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
