package main

import (
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/migrations"
	"github.com/kevinmranda/GoAPI/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrations.SyncDatabase()
}

func main() {
	routesRoutes()
}
