package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_name string
	last_name  string
	password   string
	gender     string
	birthdate  time.Time
	address    string
	email      string `gorm:"unique"`
	mobile     string
}
