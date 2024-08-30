package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerEmail string
	Cart          Cart    // One-to-One relationship with Cart
	Orders        []Order // One-to-Many relationship with Order
}
