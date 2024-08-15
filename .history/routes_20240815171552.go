package main

func routes () {
	r := gin.Default()


	

	r.Run()
	r.POST("/createAccount", controllers.CreateAccount)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/insertPhoto", controllers.AddPhoto)
}