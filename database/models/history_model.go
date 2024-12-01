package models

import (
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	ESPcode string
	KwH     float64
	Price   float64
}
