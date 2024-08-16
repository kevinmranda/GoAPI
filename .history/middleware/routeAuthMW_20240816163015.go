package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func RouteAuthMW(c *gin.Context) {

	// chaeck if user role is admin
	var user models.User
	result := initializers.DB.Where("name = ?", "jinzhu").First(&user)
	if result.Error != nil || user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Attach user to context
	c.Set("user", user)

	// Continue to the next handler
	c.Next()

}
