package router

import (
	"com.ocbc.smb/constants"
	"com.ocbc.smb/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))

	AccountController := new(controller.AccountController)

	router_version := ""

	router_version = "/api/" + constants.API_VERSION1 + "/"
	api := r.Group(router_version + "account")
	api.GET("", AccountController.GetUser)
	api.POST("", AccountController.GetUser)

	return r
}
