package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title string
	Description Text
	Filename String //(path to the high-quality image)
	// low_res_filename: String (path to the low-quality watermarked image)
	Price Decimal
	Uploaded_by uint
}
