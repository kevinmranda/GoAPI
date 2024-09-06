package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	// Customer_email string //(Email for the order, used for sending download links)
	Total_amount   float64
	Status         string  `gorm:"default:pending"` //(enum: "pending", "completed", "canceled")
	Photos         []Photo `gorm:"many2many:order_photos;"`
	Payment        Payment // One-to-One relationship with Payment
	Customers     [uint]
}
