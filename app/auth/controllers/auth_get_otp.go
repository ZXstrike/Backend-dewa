package controllers

import (
	"time"
	"zxsttm/database"
	"zxsttm/database/models"
	"zxsttm/pkg/mail"
	"zxsttm/pkg/otp"

	"github.com/gin-gonic/gin"
)

func GetAuthGetOtp(context *gin.Context) {

	db := database.DB

	email := context.Query("email")

	var otpString string

	var userOtp models.UserOTP

	if err := db.Where("email = ?", email).First(&userOtp).Error; err != nil {
		otpString = otp.GenerateOTP()
		userOtp.Email = email
		userOtp.OTP = otpString
		t := time.Now().Add(5 * time.Minute)
		userOtp.TimeOut = &t
		db.Create(&userOtp)
	} else {
		otpString = userOtp.OTP
		t := time.Now().Add(5 * time.Minute)
		userOtp.TimeOut = &t
		db.Save(&userOtp)
	}

	err := mail.SendMail([]string{email}, []string{}, "OTP", otpString)

	if err != nil {
		context.JSON(500, gin.H{
			"error": "Failed to send OTP",
		})
		return
	}

	context.JSON(200, gin.H{
		"message": "Auth check",
	})

}
