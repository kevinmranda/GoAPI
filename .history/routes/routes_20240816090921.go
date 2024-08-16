package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/controllers"
	"github.com/kevinmranda/GoAPI/middleware"
)

//all web routes defined here
func Routes() {
	r := gin.Default()

	//User Routes
	r.POST("/createAccount", controllers.CreateAccount)
	r.DELETE("/deleteUser/:id", controllers.DeleteUser)
	r.GET("/getUser/:id", controllers.GetUser)
	r.GET("/getUsers/", controllers.GetUsers)
	r.PUT("/updateUser/:id", controllers.UpdateUser)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	//Photo Routes
	r.POST("/insertPhoto", controllers.AddPhoto)
	r.DELETE("/deletePhoto/:id", controllers.DeletePhoto)
	r.GET("/getPhoto/:id", controllers.GetPhoto)
	r.GET("/getPhotos/", controllers.GetPhotos)
	r.PUT("/updatePhoto/:id", controllers.UpdatePhoto)

	//Order Routes
	r.POST("/addOrder", controllers.AddOrder())
	r.DELETE("/removeOrder/:id", controllers.RemoveOrder())
	r.GET("/getOrder/:id", controllers.GetOrder
	r.GET("/getOrders/", controllers.GetPhotos)
	r.PUT("/updateOrder/:id", controllers.UpdatePhoto)

	//Payment Routes

	r.Run()

}
