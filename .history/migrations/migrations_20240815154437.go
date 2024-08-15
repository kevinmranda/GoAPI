package migrations

import (
	"github.com/kevinmranda/GoAPI/initializers"
	"github.com/kevinmranda/GoAPI/models"
)

func SyncDatabase() {
	//Migrate Schema
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Photo{})
	initializers.DB.AutoMigrate(&models.Order{})
	initializers.DB.AutoMigrate(&models.User{})
}
