package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Customer_email string //(Email for the order, used for sending download links)
	Total_amount   float64
	Status         string //(enum: "pending", "completed", "canceled")
	Photos []Photo `gorm:"many2many:orderPhoto;"`
}
