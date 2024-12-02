package controllers

import (
	"zxsttm/pkg/otp"

	"github.com/gin-gonic/gin"
)

func AuthOtp(context *gin.Context) {
	email := context.Query("email")
	input_otp := context.Query("otp")

	if otp.CheckOTP(email, input_otp) {
		context.JSON(200, gin.H{
			"message": "OTP verified",
		})
		return
	}

	context.JSON(400, gin.H{
		"error": "OTP salah or sudah berakhir",
	})
}
