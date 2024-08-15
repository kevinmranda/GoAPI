package controllers

import "github.com/gin-gonic/gin"

func AddPhoto(c *gin.Context) {
	//struct of photo
	var body struct {
		Title       string 
		Description string
		Filename    string //(path to the high-quality image)
		Price       float64
		User_id     uint //Uploaded by
	}

	//get input from json

	//bind input to struct

	//continue
}
