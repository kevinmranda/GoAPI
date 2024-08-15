package controllers

import "github.com/gin-gonic/gin"

func AddPhoto(c *gin.Context) {
	//struct of the request body
	var body struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Filename    string `json:"filename" binding:"required"`//(path to the high-quality image)
		Price       float64 `json:"price" binding:"required"`
		User_id     uint `json:"user_id" binding:"required"`//Uploaded by
	}

	//get input from json


	//bind input to struct

	//continue
}
