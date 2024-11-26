package controllers

import (
	"net/http"
	"strconv"

	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func GetESPByID(c *gin.Context) {
	db := database.DB
	id := c.Param("id")
	espID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var esp models.ESP
	if err := db.Preload("DHTs").Preload("PZEMs").First(&esp, espID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ESP not found"})
		return
	}

	// if len(esp.DHTs) == 0 && len(esp.PZEMs) == 0 {
	// 	c.JSON(http.StatusOK, gin.H{"message": "Sensor not found"})
	// 	return
	// }

	c.JSON(http.StatusOK, esp)
}
