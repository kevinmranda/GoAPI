package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Order_id       uint //(Foreign Key to Order)
	Amount         float64
	status         string //(enum: "pending", "completed", "failed")
	payment_method string //(e.g., "credit_card", "paypal")
	transaction_id string //(Identifier from payment gateway)
}
