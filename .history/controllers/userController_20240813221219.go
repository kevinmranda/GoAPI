package controllers

import "github.com/gin-gonic/gin"

func CreateAccount(c *gin.Context) {
	//get contents from body of request
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Body",
		})
		return
	}

	//hash password

	//create user

	//respond 
}

func Login(c *gin.Context) {

}
