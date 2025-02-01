package main

import (
	"nexcommerce/middlewares"
	"nexcommerce/models"
	"nexcommerce/responses"
	"nexcommerce/routes/auth"
	v1 "nexcommerce/routes/v1"
	"nexcommerce/stores"
	"nexcommerce/utils/config"
	"nexcommerce/utils/logger"

	"github.com/gin-gonic/gin"
)

func init() {
	logger.SetupLogger()
	config.LoadConfig()

}

func main() {
	// setup gin server mode based on yaml config
	switch config.Configs.Server.Mode {
	case "release":
		logger.Logger.Println("Starting server in release mode")
		gin.SetMode(gin.ReleaseMode)
	default:
		logger.Logger.Println("Starting server in debug mode")
		gin.SetMode(gin.DebugMode)
	}

	stores.GetDb().AutoMigrate(&models.User{})

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	// health endpoint
	router.GET("/health", func(c *gin.Context) {
		responses.Ok(c, nil)
	})

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// add all of the routes
	v1.UserRoutes(router.Group(v1.RouteGroupName))
	v1.ApiDocRoutes(router.Group(v1.RouteGroupName))
	auth.Routes(router.Group(auth.RouteGroupName))

	router.Run(config.Configs.Server.Port)
}
