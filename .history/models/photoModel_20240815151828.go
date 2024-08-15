package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	title: String
	description: Text
	filename: String //(path to the high-quality image)
	// low_res_filename: String (path to the low-quality watermarked image)
	price: Decimal
	uploaded_by
}
