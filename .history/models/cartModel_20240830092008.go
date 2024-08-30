package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Photos   []Photo
	Customer Customer
}
