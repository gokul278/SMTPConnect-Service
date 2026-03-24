package authenticationroutes

import (
	authenticationcontrollers "mailstitch/controller/Authentication"

	"github.com/gin-gonic/gin"
)

func InitAuthenticationRoutes(router *gin.RouterGroup) {
	route := router.Group("/authentication")
	route.POST("/login", authenticationcontrollers.SignInController())
	route.POST("/signup", authenticationcontrollers.SignUpController())
}
