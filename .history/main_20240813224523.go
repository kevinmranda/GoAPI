package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/controllers"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/migrat"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrations.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/createAccount", controllers.CreateAccount)

	r.Run()
}
