package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/controllers"
	"github.com/kevinmranda/GoAPI/middleware"
)

// all web routes defined here
func Routes() {
	r := gin.Default()

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
