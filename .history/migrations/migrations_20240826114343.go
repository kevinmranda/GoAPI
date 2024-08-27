package migrations

import (
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

// migrations for the database are found here
func SyncDatabase() {
	//Migrate Schema
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Photo{})
	initializers.DB.AutoMigrate(&models.Order{})
	initializers.DB.AutoMigrate(&models.Payment{})
	initializers.DB.AutoMigrate(&models.Token{})
	initializers.DB.AutoMigrate(&models.ActivityLog{})
}
