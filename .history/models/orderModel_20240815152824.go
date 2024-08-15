package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Customer_email string //(Email for the order, used for sending download links)
tTotal_amount float64
status string //(enum: "pending", "completed", "canceled")
}





customer_email: String (Email for the order, used for sending download links)
total_amount: Decimal
status: String (enum: "pending", "completed", "canceled")