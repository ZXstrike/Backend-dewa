package relay

import (
	"net/http"
	"strconv"

	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func GetRelayState(c *gin.Context) {
	db := database.DB
	code := c.Param("espcode")
	espCode, err := strconv.ParseUint(code, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	esp := models.ESP{}

	if err := db.First(&esp, "espcode = ?", espCode).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ESP tidak ditemukan"})
		return
	}

	var relay models.Relay
	if err := db.First(&relay, "es_pid = ?", esp.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Relay tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": relay})
}
