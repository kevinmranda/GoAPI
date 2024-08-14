package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	first_name string
	last_name  string
	password   string
	gender     string
	birthdate  time.Time
	address    string
	email      string
	mobile     string
}
