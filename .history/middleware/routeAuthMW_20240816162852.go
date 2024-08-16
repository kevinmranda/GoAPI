package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func RouteAuthMW(c *gin.Context) {

// chaeck if user role is admin
		var user models.User
		result := initializers.DB.First(&user)
		if result.Error != nil || user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach user to context
		c.Set("user", user)

		// Continue to the next handler
		c.Next()
	
		c.AbortWithStatus(http.StatusUnauthorized)
	
}
