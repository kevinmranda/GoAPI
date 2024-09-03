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
	r.GET("/getPhoto/:filename", controllers.GetPhoto)
	r.GET("/getPhotos/:id", controllers.GetPhotos)
	r.GET("/getAllPhotos/", controllers.GetAllPhotos)

	// Order Routes
	r.POST("/addOrder/", controllers.AddOrder)
	r.GET("/getOrder/:id", controllers.GetOrder)
	r.GET("/getOrders/:id", controllers.GetOrders)

	// Payment Routes
	r.POST("/payOrder", controllers.AddPayment)
	r.GET("/getPayment/:id", controllers.GetPayment)
	r.GET("/getPayments/:id", controllers.GetPayments)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware,
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:4200"}, // Allow frontend's origin
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}), middleware.LogRequestResponseMiddleware())
	{
		// User Routes
		protected.DELETE("/deleteUser/:id", controllers.DeleteUser)
		protected.PUT("/updateUser/:id", controllers.UpdateUser)
		protected.PUT("/updateUserPassword/:id", controllers.UpdateUserPassword)
		protected.PUT("/updateUserPreferences/:id", controllers.UpdateUserPreferences)
		protected.GET("/userPreferences/:id", controllers.GetUserPreferences)

		// Photo Routes
		protected.POST("/upload/", controllers.Upload)
		protected.POST("/insertPhoto/:id", controllers.AddPhoto)
		protected.DELETE("/deletePhoto/:id", controllers.DeletePhoto)
		protected.PUT("/updatePhoto/:id", controllers.UpdatePhoto)

		// Order Routes
		protected.DELETE("/removeOrder/:id", controllers.RemoveOrder)
		protected.PUT("/updateOrder/:id", controllers.UpdateOrder)

		// Payment Routes
		protected.DELETE("/deletePayment/:id", controllers.DeletePayment)
		protected.PUT("/updatePayment/:id", controllers.UpdatePayment)

		//Customer Routes
		protected.POST("/customerLogin/", controllers.CustomerAuthentication)
		protected.POST("/customerJoin/", controllers.AddCustomer)

		//Logs Routes
		protected.GET("/logs", controllers.GetLogs)

	}

	r.Run()

}
