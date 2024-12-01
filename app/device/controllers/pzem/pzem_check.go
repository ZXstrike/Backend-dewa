package pzem

import (
	"net/http"
	"strconv"

	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func GetPZEMByID(c *gin.Context) {
	db := database.DB
	id := c.Param("id")
	pzemID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var pzem models.PZEM
	if err := db.First(&pzem, pzemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PZEM not found"})
		return
	}

	c.JSON(http.StatusOK, pzem)
}
