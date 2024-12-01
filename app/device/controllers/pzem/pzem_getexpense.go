package pzem

import (
	"strconv"
	"zxsttm/database/models"

	"zxsttm/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetExpense(c *gin.Context) {
	db := database.DB

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid page number",
		})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid page size",
		})
		return
	}

	offset := (page - 1) * pageSize

	var projects []models.Monthly
	db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}).Find(&projects)

	c.JSON(200, gin.H{
		"message":  "Get all projects",
		"projects": projects,
	})
}
