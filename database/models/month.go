package models

import (
	"gorm.io/gorm"
)

type Monthly struct {
	gorm.Model
	TotalKwH float64
	Perbandingan float64
	Expense float64
	ESPid uint
}