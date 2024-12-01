package esp

import (
	"net/http"

	"time"
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func EspCheck(c *gin.Context) {
	db := database.DB
	id := c.Query("espcode")

	var esp models.ESP
	if err := db.Where("es_pcode = ?", id).First(&esp).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ESP not found", "id": id})
		return
	}

	esp.IsActive = true

	t := time.Now().Add(30 * time.Second)
	esp.LiveTimeOut = &t

	db.Save(&esp)

	c.JSON(http.StatusOK, esp)
}
