package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/controllers"
)

// all web routes defined here
func Routes() {
	r := gin.Default()

	// CORS Middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Allow your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// User Routes
	r.POST("/createAccount", controllers.CreateAccount)
	r.POST("/login", controllers.Login)
	r.GET("/getUser/:id", controllers.GetUser)
	r.GET("/getUsers/", controllers.GetUsers)

	// Photo Routes
	r.GET("/getPhoto/:id", controllers.GetPhoto)
	r.GET("/getPhotos/", controllers.GetPhotos)

	// Order Routes
	r.POST("/addOrder", controllers.AddOrder)
	r.GET("/getOrder/:id", controllers.GetOrder)
	r.GET("/getOrders/", controllers.GetOrders)

	// Payment Routes
	r.POST("/payOrder", controllers.AddPayment)
	r.GET("/getPayment/:id", controllers.GetPayment)
	r.GET("/getPayments/", controllers.GetPayments)

	protected := r.Group("/")
	// // protected.Use(middleware.AuthMiddleware)
	{
		// User Routes
		protected.DELETE("/deleteUser/:id", controllers.DeleteUser)
		protected.PUT("/updateUser/:id", controllers.UpdateUser)
		protected.POST("/sendResetPasswordEmail/:id", controllers.SendRestPasswordEmail)

		// Photo Routes
		protected.POST("/insertPhoto", controllers.AddPhoto)
		protected.DELETE("/deletePhoto/:id", controllers.DeletePhoto)
		protected.PUT("/updatePhoto/:id", controllers.UpdatePhoto)

		// Order Routes
		protected.DELETE("/removeOrder/:id", controllers.RemoveOrder)
		protected.PUT("/updateOrder/:id", controllers.UpdateOrder)

		// Payment Routes
		protected.DELETE("/deleteOrder/:id", controllers.DeletePayment)
		protected.PUT("/updatePayment/:id", controllers.UpdatePayment)
	}

	r.Run()
}
