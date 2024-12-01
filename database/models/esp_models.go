package models

import (
	"time"

	"gorm.io/gorm"
)

type ESP struct {
	gorm.Model
	ESPcode     string `gorm:"unique"`
	IsActive    bool
	LiveTimeOut *time.Time `gorm:"default:null"`
}
