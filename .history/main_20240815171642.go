package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinmranda/GoAPI/controllers"
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/middleware"
	"github.com/kevinmranda/GoAPI/migrations"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrations.SyncDatabase()
}

func main() {
	route()
}
