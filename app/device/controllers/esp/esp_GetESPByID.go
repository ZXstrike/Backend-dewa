package esp

import (
	"net/http"

	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func GetESPByID(c *gin.Context) {
	db := database.DB
	id := c.Query("espcode")

	var esp models.ESP
	if err := db.Where("es_pcode = ?", id).First(&esp).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ESP not found", "id": id})
		return
	}

	// if len(esp.DHTs) == 0 && len(esp.PZEMs) == 0 {
	// 	c.JSON(http.StatusOK, gin.H{"message": "Sensor not found"})
	// 	return
	// }

	c.JSON(http.StatusOK, esp)
}
