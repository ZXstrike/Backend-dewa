package routes

import (
	"zxsttm/app/device/controllers"

	"github.com/gin-gonic/gin"
)

func DeviceRoutes(router *gin.Engine) {
	deviceRoutes := router.Group("devices")

	deviceRoutes.GET("/check", controllers.GetDHTByID)
	deviceRoutes.POST("/add", controllers.DHTadd)
}
