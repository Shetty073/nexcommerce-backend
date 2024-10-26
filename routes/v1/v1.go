package v1

import (
	"net/http"
	"nexcommerce/middlewares"

	"github.com/gin-gonic/gin"
)

func PlaceHolder(c *gin.Context) {}

func Routes(router *gin.RouterGroup) {
	router.Use(middlewares.JWTMiddleware())
	{
		router.Handle(http.MethodPost, "/user/create", PlaceHolder)
		router.Handle(http.MethodPost, "/user/read", PlaceHolder)
		router.Handle(http.MethodPost, "/user/read/:id", PlaceHolder)
		router.Handle(http.MethodPost, "/user/update/:id", PlaceHolder)
		router.Handle(http.MethodPost, "/user/delete/:id", PlaceHolder)
	}

	return
}
