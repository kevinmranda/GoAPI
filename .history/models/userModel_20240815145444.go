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
	UserGroup
}
