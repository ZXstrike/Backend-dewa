package relay

import (
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func RelayUpdateState(c *gin.Context) {
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

	var relay models.Relay

	if err := db.Where("espid = ?", body.ESPID).First(&relay).Error; err != nil {
		relay = models.Relay{
			ESPID:     uint(body.ESPID),
			Condition: body.Condition,
		}
		db.Create(&relay)
	} else {
		relay.Condition = body.Condition
		db.Save(&relay)
	}

	c.JSON(200, gin.H{"message": "Update relay state success"})
}
