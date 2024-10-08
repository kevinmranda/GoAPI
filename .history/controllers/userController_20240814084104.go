package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
	"golang.org/x/crypto/bcrypt"
)

func validateUserInput(firstName, lastName, email string, ) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := 
	parts := strings.Split(email, "@")
	return len(strings.Split(email, "@")) == 2 && strings.Contains(strings.Split(email, "@")[1], ".") && len(strings.Split(parts[1], ".")) == 2


	return isValidName, isValidEmail
}

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
	hash, err := bcrypt.GenerateFromPassword([]byte(body.password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//create user
	user := models.User{
		First_name: body.first_name,
		Last_name:  body.last_name,
		Password:   string(hash),
		Gender:     body.gender,
		Birthdate:  body.birthdate,
		Address:    body.address,
		Email:      body.email,
		Mobile:     body.mobile,
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//respond
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {

}
