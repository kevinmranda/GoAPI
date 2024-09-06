package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerEmail string
	Orders        []Order `gorm:"many2many:order_customers;"`
}
