package controllers

import (
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func DHTadd(c *gin.Context) {
	db := database.DB

	var body struct {
		ESPID       string  `json:"ESPid" binding:"required"`
		Temperature float64 `json:"Temperature" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	Convert ESPID from string to uint
	espidUint, err := strconv.ParseUint(body.ESPID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid ESPID format",
		})
		return
	}

	dht := models.DHT{
		ESPID:       uint(espidUint), // now assigning as uint
		Temperature: body.Temperature,
	}

	// Save the dht record to the database
	if err := db.Create(&dht).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create record"})
		return
	}

	c.JSON(200, gin.H{"message": "Record created successfully"})
}
