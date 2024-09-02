// Photo Model
package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title       string
	Description string
	Filename    string `gorm:"unique"` //(path to the high-quality image)
	Price       float64
	User_id     uint    //Uploaded by
	Orders      []Order `gorm:"many2many:order_photos;"`
}
