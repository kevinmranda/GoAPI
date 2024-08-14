package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	//get contents from body of request
	var body struct {
		first_name string
		last_name  string
		password   string
		gender     string
		birthdate  time.Time
		address    string
		email      string
		mobile     string
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
