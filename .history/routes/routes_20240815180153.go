package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/controllers"
	"github.com/kevinmranda/GoAPI/middleware"
)

//all web routes defined here
func Routes() {
	r := gin.Default()

	r.POST("/createAccount", controllers.CreateAccount)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/insertPhoto", controllers.AddPhoto)
	r.DELETE("/deletePhoto/:id", controllers.DeletePhoto)

	r.Run()

}
