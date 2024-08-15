package controllers

import (
	"net/http"
	"os"
	"regexp"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(c *gin.Context) {
	// Get contents from body of request
	var body struct {
		First_name string `json:"first_name" binding:"required"`
		Last_name  string `json:"last_name" binding:"required"`
		Password   string `json:"password" binding:"required"`
		Gender     string `json:"gender" binding:"required"`
		Birthdate  string `json:"birthdate" binding:"required"`
		Address    string `json:"address" binding:"required"`
		Email      string `json:"email" binding:"required"`
		Mobile     string `json:"mobile" binding:"required"`
	}

	// Bind JSON input to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Parse the birthdate string into a time.Time object
	birthdate, err := time.Parse("2006-01-02", body.Birthdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid birthdate format, expected YYYY-MM-DD",
		})
		return
	}

	// Validate user input
	isValidName, isValidEmail, isValidPassword := validateUserInput(
		body.First_name,
		body.Last_name,
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
			Password:   string(hash),
			Gender:     body.Gender,
			Birthdate:  birthdate,
			Address:    body.Address,
			Email:      body.Email,
			Mobile:     body.Mobile,
		}
		result := initializers.DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"id":    "2000",
				"error": "Failed to create user",
			})
			return
		}

		// Respond
		c.JSON(http.StatusOK, gin.H{
			"id":      2001,
			"message": "User created successfully",
			"data":    user,
		})
	} else {
		if !isValidName {
			c.JSON(http.StatusBadRequest, gin.H{
				"id" : 2002,
				"error": "Invalid first name or last name format",
			})
		}
		if !isValidEmail {
			c.JSON(http.StatusBadRequest, gin.H{
				"id" : 2003,
				"error": "Invalid email format",
			})
		}
		if !isValidPassword {
			c.JSON(http.StatusBadRequest, gin.H{
				"id" : 2004,
				"error": "Weak password. The password should be at least 8 characters long and include special characters.",
			})
		}
	}
}

func Login(c *gin.Context) {
	//get email and password off body
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON input to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//look up for user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"id" : 2005,
			"error": "Invalid email",
			"data" : "",
		})
		return
	}

	//compare hash and password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id" : 2006,
			"error": "Invalid password",
			"data" : "",
		})
		return
	}

	//create token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(), //1 hour expiry
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":2007,
			"error": "Failed to create Token",
		})
		return
	}

	//make and set cookie
	SetAuthCookie(c, tokenString)

	//login
	c.JSON(200, gin.H{
		"id":      2008,
		"message": "Login Successfully",
		"data":    user,
		"token":
	})
}

// validate user with token
func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"id":      2009,
		"message": "Success",
		"user":    user,
	})
}

// validates user inputs
func validateUserInput(firstName, lastName, email, password string) (bool, bool, bool) {
	isValidName := len(firstName) > 0 && len(lastName) > 0
	isValidEmail := validateEmail(email)
	isValidPassword := validatePassword(password)

	return isValidName, isValidEmail, isValidPassword
}

// validateEmail checks if the email is in a valid format.
func validateEmail(email string) bool {
	// Simple regex pattern for email validation
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

// validatePassword checks if the password is strong enough
func validatePassword(password string) bool {
	var hasMinLen, hasUpper, hasLower, hasNumber, hasSpecial bool
	hasMinLen = len(password) >= 8

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// set cookie
func SetAuthCookie(c *gin.Context, tokenString string) {
	// Check if the "Auth" cookie is present
	if cookie, err := c.Cookie("Auth"); err == nil && cookie != "" {
		// Clear the existing "Auth" cookie
		c.SetCookie("Auth", "", -1, "/", "", false, true)
	}

	// Set the SameSite attribute
	c.SetSameSite(http.SameSiteLaxMode)

	// Set the new "Auth" cookie with a 30-day expiration
	c.SetCookie("Auth", tokenString, 3600, "/", "", false, true)
}
