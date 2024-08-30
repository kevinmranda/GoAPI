package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func AddItems(c *gin.Context) {
	// Bind JSON input to the struct
	var body struct {
		Customer_email string `json:"customer_email" binding:"required"`
		PhotoIDs       []uint `json:"photo_ids" binding:"required"`
	}

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

	cart := models.Cart{
		Customer_email: body.Customer_email,
		Photos:         photos,
	}

	if err := initializers.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    2000,
			"error": "Failed to insert the record",
		})
		return
	}
}

func GetItems(c *gin.Context) {

}

func ClearCart(c *gin.Context) {

}

func RemoveItem(c *gin.Context) {

}
