package pzem

import (
	"net/http"

	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func GetPZEMByID(c *gin.Context) {
	db := database.DB
	espcode := c.Query("id")

	var esp models.ESP

	if err := db.First(&esp, "espcode = ?", espcode).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ESP not found"})
		return
	}

	var pzem models.PZEM

	if err := db.First(&pzem, "esp_id = ?", esp.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PZEM not found"})
		return
	}

	c.JSON(http.StatusOK, pzem)
}
