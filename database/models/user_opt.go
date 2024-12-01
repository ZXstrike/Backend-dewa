package models

import (
	"time"

	"gorm.io/gorm"
)

type UserOTP struct {
	gorm.Model
	Email      string     `gorm:"unique"`
	OTP        string     `gorm:"unique"`
	TimeOut    *time.Time `gorm:"default:null"`
	IsVerified bool       `gorm:"default:false"`
}
