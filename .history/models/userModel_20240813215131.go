package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	first_name string
	last_name  string
	password   string
	gender     string
	birthdate  time.
	address    string
	email      string
	mobile     string
}
