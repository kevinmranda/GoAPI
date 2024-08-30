package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
}

func GetItems(c *gin.Context) {

}

func ClearCart(c *gin.Context) {

}

func RemoveItem(c *gin.Context) {

}
