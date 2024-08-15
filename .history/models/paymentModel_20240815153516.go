package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	order_id UUID (Foreign Key to Order)
amount Decimal
status String (enum: "pending", "completed", "failed")
payment_method string (e.g., "credit_card", "paypal")
transaction_id string (Identifier from payment gateway)
}
