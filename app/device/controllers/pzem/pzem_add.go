package pzem

import (
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func PZEMadd(c *gin.Context) {
	db := database.DB

	var body struct {
		ESPID   int     `json:"ESPid" binding:"required"`
		Current float64 `json:"Current" binding:"required"`
		Power   float64 `json:"Power" binding:"required"`
		Voltage float64 `json:"Voltage" binding:"required"`
		Energy  float64 `json:"Energy" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	var esp models.ESP
	if err := db.First(&esp, body.ESPID).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "ESP Not Found",
		})
		return
	}

	PZEM := models.PZEM{
		ESPID:   uint(body.ESPID),
		Current: float64(body.Current),
		Power:   float64(body.Power),
		Voltage: float64(body.Voltage),
		Energy:  float64(body.Energy),
	}

	if err := db.Create(&PZEM).Error; err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "Record created successfully"})
}
