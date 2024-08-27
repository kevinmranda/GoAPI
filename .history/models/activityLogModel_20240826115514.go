package models

import (
	"gorm.io/gorm"
)

type ActivityLog struct {
	gorm.Model
	Level   string
	Message string
	Details map[string]interface{} `gorm:"type:jsonb"`
}
