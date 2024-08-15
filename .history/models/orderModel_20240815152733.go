package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Title       string
	Description string
	Filename    string //(path to the high-quality image)
	// low_res_filename: String (path to the low-quality watermarked image)
	Price       float64
	Uploaded_by uint //User ID
}





customer_email: String (Email for the order, used for sending download links)
total_amount: Decimal
status: String (enum: "pending", "completed", "canceled")