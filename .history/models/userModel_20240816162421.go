package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_name string
	Last_name  string
	Password   string
	Gender     string
	Birthdate  time.Time
	Address    string
	Email      string `gorm:"unique"`
	Mobile     string
	Role     string `gorm:"default:cat"`
	Photos     []Photo // One-to-Many relationship with Photo
}
