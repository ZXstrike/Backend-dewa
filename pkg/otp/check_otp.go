package otp

import (
	"time"
	"zxsttm/database"
	"zxsttm/database/models"
)

func CheckOTP(email string, input_otp string) bool {
	db := database.DB

	var user models.User

    if err := db.Where("email = ?", email).First(&user).Error; err != nil {
        return false
    }

	var userOtp models.UserOTP

	db.Raw("SELECT * FROM user_otps WHERE email = ?", email).Scan(&userOtp)

	if userOtp.IsVerified {
		return true
	}

	// Check timeout
	current_time := time.Now()
	if userOtp.TimeOut.Before(current_time) {
		return false
	}

	if userOtp.OTP == input_otp {
		userOtp.IsVerified = true
		db.Save(&userOtp)
		return true
	}
	return false
}
