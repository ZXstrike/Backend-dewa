package models

import (
	"gorm.io/gorm"
)

type PZEM struct {
	gorm.Model 
	ESPID   uint
	Current float64
	Power   float64
	Voltage float64
	Energy  float64 
}
