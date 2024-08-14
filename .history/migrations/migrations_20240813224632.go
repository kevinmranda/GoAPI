package migrations

import (
	"github.com/kevinmranda/GoAPI/models"
)

func SyncDatabase() {
	//Migrate Schema
	initializers.DB.AutoMigrate(&models.User{})
}
