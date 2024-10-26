package main

import (
	responses "nexcommerce/common"
	"nexcommerce/middlewares"
	"nexcommerce/models"
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
		return
	})

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// add v1 routes group to the router
	v1.Routes(router.Group("/v1"))
	auth.Routes(router.Group("/auth"))
	router.Run(config.Configs.Server.Port)
}
