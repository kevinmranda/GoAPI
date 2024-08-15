package main

func main() {
	r.POST("/createAccount", controllers.CreateAccount)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/insertPhoto", controllers.AddPhoto)
}