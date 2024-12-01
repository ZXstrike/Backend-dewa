package dht

import (
	"net/http"
	"strconv"

	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func GetDHTByID(c *gin.Context) {
	db := database.DB
	id := c.Query("ESPid")
	dhtID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": id})
		return
	}

	var dht models.DHT
	if err := db.Where("esp_id = ?", dhtID).First(&dht).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sensor tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, dht)
}
