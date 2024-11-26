package controllers

import (
	"net/http"
	"strconv"

	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func GetDHTByID(c *gin.Context) {
	db := database.DB
	id := c.Query("id")
	dhtID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var dht models.DHT
	if err := db.First(&dht, dhtID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sensor tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, dht)
}
