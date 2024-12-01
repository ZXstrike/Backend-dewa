package models

import (
	"gorm.io/gorm"
)

type Relay struct {
	gorm.Model
	ESPID     uint
	Condition bool
}
