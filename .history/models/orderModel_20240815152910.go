package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Customer_email string //(Email for the order, used for sending download links)
	Total_amount   float64
	status         string  //(enum: "pending", "completed", "canceled")
}
