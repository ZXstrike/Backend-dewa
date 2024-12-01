package relay

import (
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func RELAYadd(c *gin.Context) {
	db := database.DB

	var body struct {
		ESPID     int  `json:"ESPid" binding:"required"`
		Condition bool `json:"Condition" binding:"required"`
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

	Relay := models.Relay{
		ESPID:     uint(body.ESPID),
		Condition: bool(body.Condition),
	}

	if err := db.Create(&Relay).Error; err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "Record created successfully"})
}