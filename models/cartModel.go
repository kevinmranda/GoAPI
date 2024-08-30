package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Photos     []Photo `gorm:"many2many:cart_photos;"`
	CustomerID uint
}
