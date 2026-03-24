package mailroutes

import (
	mailcontrollers "mailstitch/controller/Mail"
	accesstoken "mailstitch/internal/Helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func InitMailRoutes(api *gin.RouterGroup) {
	route := api.Group("/mail")
	route.Use(accesstoken.JWTMiddleware())
	{
		route.POST("/send", mailcontrollers.SendMailController())
		route.GET("/history", mailcontrollers.GetMailHistoryController())
	}
}
