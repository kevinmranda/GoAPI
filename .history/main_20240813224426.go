package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/controllers"
	"github.com/kevinmranda/GoAPI/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrations.Sync
}

func main() {
	r := gin.Default()

	r.POST("/createAccount", controllers.CreateAccount)

	r.Run()
}
