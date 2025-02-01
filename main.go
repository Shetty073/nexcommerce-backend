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
	"nexcommerce/utils/migrations"
	"nexcommerce/utils/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func init() {
	// Setup logger and load config
	logger.SetupLogger()
	config.LoadConfig()
	migrations.RegisterAllModels()
}

func main() {
	// Setup gin server mode based on YAML config
	switch config.Configs.Server.Mode {
	case "release":
		logger.Logger.Println("Starting server in release mode")
		gin.SetMode(gin.ReleaseMode)
	default:
		logger.Logger.Println("Starting server in debug mode")
		gin.SetMode(gin.DebugMode)
	}

	// Initialize database and auto-migrate models
	stores.GetDb().AutoMigrate(&models.User{})

	// Set up the Gin router with default settings
	router := gin.Default()

	// Use CORS middleware for cross-origin requests
	router.Use(middlewares.CORSMiddleware())

	// Define the health check route
	router.GET("/health", func(c *gin.Context) {
		responses.Ok(c, nil)
	})

	// Set trusted proxies (if needed)
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Register routes
	auth.Routes(router.Group(auth.RouteGroupName))
	v1.UserRoutes(router.Group(v1.RouteGroupName))
	v1.ApiDocRoutes(router.Group(v1.RouteGroupName))

	// Register custom validator
	validate := validator.New()
	validators.RegisterValidators(validate)

	// Run the Gin server
	router.Run(config.Configs.Server.Port)
}
