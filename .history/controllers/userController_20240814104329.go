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
	// Get contents from body of request
	var body struct {
		First_ame string    `json:"first_name" binding:"required"`
		Last_ame  string    `json:"last_name" binding:"required"`
		Password  string    `json:"password" binding:"required"`
		Gender    string    `json:"gender" binding:"required"`
		Birthdate time.Time `json:"birthdate" binding:"required"`
		Address   string    `json:"address" binding:"required"`
		Email     string    `json:"email" binding:"required,email"`
		Mobile    string    `json:"mobile" binding:"required"`
	}

	// Bind JSON input to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Validate user input
	isValidName, isValidEmail, isValidPassword := validateUserInput(
		body.FirstName,
		body.LastName,
		body.Email,
		body.Password,
	)

	if isValidName && isValidEmail && isValidPassword {
		// Hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash password",
			})
			return
		}

		// Create user
		user := models.User{
			First_name: body.First_name,
			Last_name:  body.Last_name,
			Password:  string(hash),
			Gender:    body.Gender,
			Birthdate: body.Birthdate,
			Address:   body.Address,
			Email:     body.Email,
			Mobile:    body.Mobile,
		}
		result := initializers.DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create user",
			})
			return
		}

		// Respond
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
				"error": "Weak password. The password should be at least 8 characters long and include special characters.",
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
