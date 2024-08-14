package main

import (
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func SyncDatabase() {
	//Migrate Schema
	initializers.DB.AutoMigrate(&models.User{})
}
