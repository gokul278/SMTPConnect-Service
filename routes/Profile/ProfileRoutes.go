package profileroutes

import (
	usercontrollers "smtpconnect/controller/User"
	accesstoken "smtpconnect/internal/Helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func InitProfileRoutes(router *gin.RouterGroup) {
	route := router.Group("/profile")
	route.GET("/user", accesstoken.JWTMiddleware(), usercontrollers.UserProfileController())
}
