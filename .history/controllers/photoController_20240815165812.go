package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPhoto(c *gin.Context) {
	//struct of the request body
	var body struct {
		Title       string  `json:"title" binding:"required"`
		Description string  `json:"description" binding:"required"`
		Filename    string  `json:"filename" binding:"required"` //(path to the high-quality image)
		Price       float64 `json:"price" binding:"required"`
		User_id     uint    `json:"user_id" binding:"required"` //Uploaded by
	}

	//Get contents from body of request and Bind JSON input to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//continue
	
}
