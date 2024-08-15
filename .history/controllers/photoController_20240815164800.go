package controllers

import "github.com/gin-gonic/gin"

func AddPhoto(c *gin.Context) {
	//struct of photo
	var body struct {
		Title       string `json:"first_name" binding:"required"`
		Description string `json:"first_name" binding:"required"`
		Filename    string `json:"first_name" binding:"required"`//(path to the high-quality image)
		Price       float64
		User_id     uint //Uploaded by
	}

	//get input from json

	//bind input to struct

	//continue
}
