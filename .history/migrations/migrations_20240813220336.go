package migrations

import "github.com/kevinmranda/GoAPI/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func SyncDatabase() {
	//Migrate Schema
	initializers.DB.AutoMigrate(&models.User{})
}
