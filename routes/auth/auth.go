package auth

import (
	"net/http"

	"nexcommerce/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.Handle(http.MethodPost, "/login", controllers.LoginController)
	router.Handle(http.MethodPost, "/register", controllers.RegisterController)

	return
}
