package main

import (
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/migrations"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrations.SyncDatabase()
}

func main() {
	outes()
}
