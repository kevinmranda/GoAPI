package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title: String
	Description: Text
	Filename: String //(path to the high-quality image)
	// low_res_filename: String (path to the low-quality watermarked image)
	price: Decimal
	uploaded_by: uint
}
