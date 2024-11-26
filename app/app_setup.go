package app

import (
	auth_routes "zxsttm/app/auth/routes"
	"zxsttm/app/device/routes"

	"github.com/gin-gonic/gin"
)

func InitApp(router *gin.Engine) {

	// add app routes here

	auth_routes.AuthRoutes(router)
	routes.DeviceRoutes(router)
}
