package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func AddItems(c *gin.Context) {
	customer_idStr := c.Param("customer_id")
	customer_id64, _ := strconv.ParseUint(customer_idStr, 10, 32)
	customer_id := uint(customer_id64)
	// Bind JSON input to the struct
	var body struct {
		PhotoIDs []uint `json:"photo_ids" binding:"required"`
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
		CustomerID: customer_id,
		Photos:     photos,
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
	customer_id := c.Param("customer_id")
	var cart []models.Cart
	result := initializers.DB.Where("customer_id = ?", customer_id).Preload("Photos").First(&cart)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not found",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":   2001,
		"cart": cart,
	})

}

func ClearCart(c *gin.Context) {
	customer_id := c.Param("customer_id")
	var cart []models.Cart
	result := initializers.DB.Where("customer_id = ?", customer_id).Delete(&cart)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "record not found",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2012,
		"message": "Cart cleared successfully",
	})
}

func RemoveItem(c *gin.Context) {
	customer_id := c.Param("customer_id")
	photo_id := c.Param("photo_id")
	var cart models.Cart

	// Find the cart by customer_id
	if err := initializers.DB.Where("customer_id = ?", customer_id).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "cart not found",
		})
		return
	}

	// Find the photo by photo_id
	var photo models.Photo
	if err := initializers.DB.Where("id = ?", photo_id).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id":    2011,
			"error": "photo not found",
		})
		return
	}

	// Delete the association between the cart and the photo
	if err := initializers.DB.Model(&cart).Association("Photos").Delete(&photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"id":    2013,
			"error": "could not remove item from cart",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"id":      2012,
		"message": "Item removed successfully",
	})
}
