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

func CreateAccount(c *gin.Context) {
	//get contents from body of request
	var body struct {
		First_name string
		Last_name  string
		Password   string
		Gender     string
		Birthdate  time.Time
		Address    string
		Email      string
		Mobile     string
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"userfn" : body.first_name,
	// 	"userln" : body.last_name,
	// 	"userp" : body.password,
	// 	"userg" : body.gender,
	// 	"userbd" : body.birthdate,
	// 	"userad" : body.address,
	// 	"userem" : body.email,
	// 	"usermob" : body.mobile,
	// })

	isValidName, isValidEmail, isValidPassword := validateUserInput(
		body.First_name,
		body.Last_name,
		body.Email,
		body.Password)

	if isValidName && isValidEmail && isValidPassword {
		//hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash password",
			})
			return
		}

		//create user
		user := models.User{
			First_name: body.First_name,
			Last_name:  body.Last_name,
			Password:   string(hash),
			Gender:     body.Gender,
			Birthdate:  body.Birthdate,
			Address:    body.Address,
			Email:      body.Email,
			Mobile:     body.Mobile,
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
	} else {
		if !isValidName {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid first name or last name",
			})
		}
		if !isValidEmail {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid email",
			})
		}
		if !isValidPassword {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Weak password, the password should be about 8 characters with special keys",
			})
		}
	}

}

func Login(c *gin.Context) {

}

func validateUserInput(firstName, lastName, email, password string) (bool, bool, bool) {
	// Validate name
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	// Validate email
	isValidEmail := len(strings.Split(email, "@")) == 2 && strings.Contains(strings.Split(email, "@")[1], ".") && len(strings.Split(strings.Split(email, "@")[1], ".")) == 2

	// Validate password
	isValidPassword := func(password string) bool {
		if len(password) <= 8 {
			return false
		}

		for _, char := range password {
			if strings.ContainsRune("!@#$%^&*()_+-./:;<=|>?", char) {
				return true
			}
		}
		return false
	}

	return isValidName, isValidEmail, isValidPassword(password)
}
