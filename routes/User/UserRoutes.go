package userroutes

import (
	usercontrollers "mailstitch/controller/User"
	accesstoken "mailstitch/internal/Helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.RouterGroup) {
	route := router.Group("/profile")
	route.GET("/user", accesstoken.JWTMiddleware(), usercontrollers.UserProfileController())
}
