package main

import (
	_ "com.ocbc.smb/database"
	"com.ocbc.smb/router"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
)

func main() {

	serverPort := beego.AppConfig.DefaultString("httpport", "8888")

	runMode := beego.AppConfig.DefaultString("gin.mode", "debug")
	gin.SetMode(runMode)

	routersInit := router.InitRouter()
	routersInit.Run(":" + serverPort)
}
