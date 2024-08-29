package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/controllers"
	"github.com/kevinmranda/GoAPI/middleware"
)

// all web routes defined here
func Routes() {
	r := gin.Default()

	// CORS Middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Allow frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}), middleware.LogRequestResponseMiddleware())

	// User Routes
	r.POST("/createAccount", controllers.CreateAccount)
	r.POST("/login", controllers.Login)
	r.GET("/getUser/:id", controllers.GetUser)
	r.GET("/getUsers/", controllers.GetUsers)
	r.POST("/sendResetPasswordEmail/", controllers.SendResetPasswordEmail)
	r.POST("/reset-password/:token", controllers.ResetPassword)

	// Photo Routes
	r.GET("/getPhoto/:id", controllers.GetPhoto)
	r.GET("/getPhotos/:id", controllers.GetPhotos)

	// Order Routes
	r.POST("/addOrder", controllers.AddOrder)
	r.GET("/getOrder/:id", controllers.GetOrder)
	r.GET("/getOrders/:id", controllers.GetOrders)

	// Payment Routes
	r.POST("/payOrder", controllers.AddPayment)
	r.GET("/getPayment/:id", controllers.GetPayment)
	r.GET("/getPayments/", controllers.GetPayments)

	// protected := r.Group("/")
	// // protected.Use(middleware.AuthMiddleware)
	// {
	// User Routes
	r.DELETE("/deleteUser/:id", controllers.DeleteUser)
	r.PUT("/updateUser/:id", controllers.UpdateUser)
	r.PUT("/updateUserPassword/:id", controllers.UpdateUserPassword)
	r.PUT("/updateUserPreferences/:id", controllers.UpdateUserPreferences)
	r.GET("/userPreferences/:id", controllers.GetUserPreferences)

	// Photo Routes
	r.POST("/upload/", controllers.Upload)
	r.POST("/insertPhoto/:id", controllers.AddPhoto)
	r.DELETE("/deletePhoto/:id", controllers.DeletePhoto)
	r.PUT("/updatePhoto/:id", controllers.UpdatePhoto)

	// Order Routes
	r.DELETE("/removeOrder/:id", controllers.RemoveOrder)
	r.PUT("/updateOrder/:id", controllers.UpdateOrder)

	// Payment Routes
	r.DELETE("/deleteOrder/:id", controllers.DeletePayment)
	r.PUT("/updatePayment/:id", controllers.UpdatePayment)

	//Logs Routes
	r.GET("/logs", controllers.GetLogs)

	// }

	r.Run()
}
