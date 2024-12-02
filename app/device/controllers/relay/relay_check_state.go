package relay

import (
	"net/http"

	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func GetRelayState(c *gin.Context) {
	db := database.DB
	espCode := c.Query("espcode")

	var esp models.ESP

	if err := db.First(&esp, "es_pcode = ?", espCode).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ESP tidak ditemukan"})
		return
	}

	var relay models.Relay
	if err := db.First(&relay, "esp_id = ?", esp.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Relay tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": relay})
}
