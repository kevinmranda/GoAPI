package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	bcrypt.GenerateFromPassword([]byte(body.password), 10)

	//create user

	//respond
}

func Login(c *gin.Context) {

}
