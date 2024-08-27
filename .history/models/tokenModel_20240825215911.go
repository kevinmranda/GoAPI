package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserID uint
	Token  string
	Time   time.Time
}
