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
	id := c.Param("id")
	relayID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var relay models.Relay
	if err := db.First(&relay, relayID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Relay tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": relay})
}
