package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title       string
	Description string
	Filename    string //(path to the high-quality image)
	// low_res_filename: String (path to the low-quality watermarked image)
	Price       float64
	Uploaded_by uint //User ID
	Orders []Order `gorm:"many2many:orderPo;"`
}
