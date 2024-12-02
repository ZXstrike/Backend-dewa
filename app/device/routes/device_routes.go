package routes

import (
	"zxsttm/app/device/controllers/dht"
	"zxsttm/app/device/controllers/esp"
	"zxsttm/app/device/controllers/pzem"
	"zxsttm/app/device/controllers/relay"

	"github.com/gin-gonic/gin"
)

func DeviceRoutes(router *gin.Engine) {
	deviceRoutes := router.Group("devices")

	deviceRoutes.GET("/check-dht", dht.GetDHTByID)
	deviceRoutes.POST("/add-dht", dht.DHTadd)
	deviceRoutes.GET("/getall-dht", dht.DhtGetAll)

	deviceRoutes.GET("/esp-getbyid", esp.GetESPByID)
	deviceRoutes.POST("/add-esp", esp.ESPadd)
	deviceRoutes.GET("/check-esp", esp.EspCheck)
	deviceRoutes.GET("/getall-esp", esp.EspGetAll)

	deviceRoutes.GET("/check-pzem", pzem.GetPZEMByID)
	deviceRoutes.POST("/add-pzem", pzem.PZEMadd)
	deviceRoutes.GET("/expense", pzem.GetExpense)
	deviceRoutes.GET("/getall-pzem", pzem.PZEMGetAll)

	deviceRoutes.GET("/get-relay-state", relay.GetRelayState)
	deviceRoutes.POST("/update-relay-state", relay.RelayUpdateState)
	deviceRoutes.GET("/getall-relay", relay.RelayGetAll)
}
