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
	r.POST("/login", controllers.Login)
	r.GET("/getUser/:id", controllers.GetUser)
	r.GET("/getUsers/", controllers.GetUsers)
	

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
    // User Routes
    protected.DELETE("/deleteUser/:id", controllers.DeleteUser)
    protected.PUT("/updateUser/:id", controllers.UpdateUser)

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
