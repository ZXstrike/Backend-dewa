package controllers

import (
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func ResetPassword(context *gin.Context) {

	db := database.DB

	var body struct {
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var user models.User

	err := db.First(&user, "email = ?", body.Email).Error

	if err != nil {
		context.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	var userOtp models.UserOTP

	err = db.First(&userOtp, "email = ?", body.Email).Error

	if err != nil {
		context.JSON(404, gin.H{
			"error": "User OTP not found",
		})
		return
	}

	if !userOtp.IsVerified {
		context.JSON(400, gin.H{
			"error": "OTP not verified",
		})
		return
	}

	otpErr := db.Delete(&userOtp).Error

	if otpErr != nil {
		context.JSON(500, gin.H{
			"error": "Error while deleting the OTP",
		})
		return
	}

	user.Password = body.Password

	userErr := db.Save(&user).Error

	if userErr != nil {
		context.JSON(500, gin.H{
			"error": "Error while updating the password",
		})
		return
	}

	context.JSON(200, gin.H{
		"message": "Reset Password",
	})
}
