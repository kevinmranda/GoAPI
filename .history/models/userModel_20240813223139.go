package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_name string
	Last_name  string
	password   string
	gender     string
	Birthdate  time.Time
	Address    string
	Email      string `gorm:"unique"`
	Mobile     string
}
