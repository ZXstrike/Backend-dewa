package esp

import (
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func ESPadd(c *gin.Context) {
	db := database.DB

	var body struct {
		ESPcode string `json:"ESPcode" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	esp := models.ESP{
		ESPcode: body.ESPcode,
	}

	if err := db.Create(&esp).Error; err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	relay := models.Relay{
		ESPID:     esp.ID,
		Condition: false,
	}

	if err := db.Create(&relay).Error; err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "Record created successfully"})
}
