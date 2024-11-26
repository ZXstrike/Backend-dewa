package models

import (
	"gorm.io/gorm"
)

type ESP struct {
	gorm.Model
	DHTs  []DHT  `gorm:"foreignKey:ESPID"`
	PZEMs []PZEM `gorm:"foreignKey:ESPID"`
	Relay []Relay
}
