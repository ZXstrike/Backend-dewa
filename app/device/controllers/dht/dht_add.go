package dht

import (
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func DHTadd(c *gin.Context) {
	db := database.DB

	var body struct {
		ESPcode       string     `json:"ESPcode" binding:"required"`
		Temperature float64 `json:"Temperature" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	var esp models.ESP
	if err := db.Where("es_pcode = ?", body.ESPcode).First(&esp).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "ESP Not Found",
		})
		return
	}

	dht := models.DHT{
		ESPID:       uint(esp.ID),
		Temperature: body.Temperature,
	}

	if err := db.Create(&dht).Error; err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "Record created successfully"})
}
