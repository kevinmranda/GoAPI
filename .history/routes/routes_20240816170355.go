package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/controllers"
	"github.com/kevinmranda/GoAPI/middleware"
)

// all web routes defined here
func Routes() {
	r := gin.Default()

	//User Routes
	r.POST("/createAccount", controllers.CreateAccount)
	r.GET("/getUser/:id", controllers.GetUser)
	r.GET("/getUsers/", controllers.GetUsers)
	r.POST("/login", controllers.Login)

	//Photo Routes
	r.GET("/getPhoto/:id", controllers.GetPhoto)
	r.GET("/getPhotos/", controllers.GetPhotos)

	//Order Routes
	r.POST("/addOrder", controllers.AddOrder)
	r.GET("/getOrder/:id", controllers.GetOrder)
	r.GET("/getOrders/", controllers.GetOrders)

	//Payment Routes
	r.POST("/payOrder", controllers.AddPayment)
	
	r.GET("/getPayment/:id", controllers.GetPayment)
	r.GET("/getPayments/", controllers.GetPayments)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware)
	{
		//User Routes
		r.DELETE("/deleteUser/:id", controllers.DeleteUser)
		r.PUT("/updateUser/:id", controllers.UpdateUser)

		//Photo Routes
		r.POST("/insertPhoto", controllers.AddPhoto)
		r.DELETE("/deletePhoto/:id", controllers.DeletePhoto)
		r.PUT("/updatePhoto/:id", controllers.UpdatePhoto)

		//Order Routes
		r.DELETE("/removeOrder/:id", controllers.RemoveOrder)
		r.PUT("/updateOrder/:id", controllers.UpdateOrder)

		//Payment Routes
		r.DELETE("/deleteOrder/:id", controllers.DeletePayment)
		r.PUT("/updatePayment/:id", controllers.UpdatePayment)
	}

	

	r.Run()

}
