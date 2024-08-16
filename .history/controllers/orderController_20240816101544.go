package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

//add order to db
func AddOrder(c *gin.Context) {
    // Bind JSON input to the struct
    var body struct {
        Customer_email string `json:"customer_email" binding:"required"`
        Total_amount   float64 `json:"total_amount" binding:"required"`
        Status         string `json:"status" binding:"required"`
        PhotoIDs       []uint `json:"photo_ids" binding:"required"`
    }

	controllers

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
