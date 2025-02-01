package v1

import (
	"net/http"
	"nexcommerce/middlewares"

	_ "nexcommerce/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var RouteGroupName = "/v1"

func PlaceHolder(c *gin.Context) {}

func UserRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.JWTMiddleware())
	{
		// User related routes
		router.Handle(http.MethodPost, "/user/create", PlaceHolder)
		router.Handle(http.MethodPost, "/user/read", PlaceHolder)
		router.Handle(http.MethodPost, "/user/read/:id", PlaceHolder)
		router.Handle(http.MethodPost, "/user/update/:id", PlaceHolder)
		router.Handle(http.MethodPost, "/user/delete/:id", PlaceHolder)
	}
}

func ApiDocRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.JWTMiddleware())
	{
		// API documentation
		router.Handle(http.MethodGet, "/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
