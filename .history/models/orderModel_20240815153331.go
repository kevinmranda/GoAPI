package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Customer_email string //(Email for the order, used for sending download links)
	Total_amount   float64
	Status         string //(enum: "pending", "completed", "canceled")
	Photo []Order `gorm:"many2many:orderPhoto;"`
}
