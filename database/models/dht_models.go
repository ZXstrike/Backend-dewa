package models

import (
	"gorm.io/gorm"
)

type DHT struct {
	gorm.Model
	// ESPID       uint
	Temperature float64
}
