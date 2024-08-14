package controllers

import (
	"net/http"
	"strings"
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

	c.JSON(http.StatusOK, gin.H{
		"user" : user,
	})

	// isValidName, isValidEmail, isValidPassword := validateUserInput(
	// 	body.first_name,
	// 	body.last_name,
	// 	body.email,
	// 	body.password)

	// if isValidName && isValidEmail && isValidPassword {
	// 	//hash password
	// 	hash, err := bcrypt.GenerateFromPassword([]byte(body.password), 10)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed to hash password",
	// 		})
	// 		return
	// 	}

	// 	//create user
	// 	user := models.User{
	// 		First_name: body.first_name,
	// 		Last_name:  body.last_name,
	// 		Password:   string(hash),
	// 		Gender:     body.gender,
	// 		Birthdate:  body.birthdate,
	// 		Address:    body.address,
	// 		Email:      body.email,
	// 		Mobile:     body.mobile,
	// 	}
	// 	result := initializers.DB.Create(&user)
	// 	if result.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed to create user",
	// 		})
	// 		return
	// 	}

	// 	//respond
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "User created successfully",
	// 	})
	// } else {
	// 	if !isValidName {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Invalid first name or last name",
	// 		})
	// 	}
	// 	if !isValidEmail {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Invalid email",
	// 		})
	// 	}
	// 	if !isValidPassword {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Weak password, the password should be about 8 characters with special keys",
	// 		})
	// 	}
	// }

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Body",
		})
		return
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
