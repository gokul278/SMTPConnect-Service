package configurationroutes

import (
	configurationcontrollers "smtpconnect/controller/Configuration"
	accesstoken "smtpconnect/internal/Helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func InitConfigurationRoutes(router *gin.RouterGroup) {
	route := router.Group("/configuration")
	route.Use(accesstoken.JWTMiddleware())
	{
		route.POST("/add", configurationcontrollers.AddConfigurationController())
		route.GET("/list", configurationcontrollers.GetConfigurationsController())
		route.POST("/update", configurationcontrollers.UpdateConfigurationController())
		route.POST("/delete", configurationcontrollers.DeleteConfigurationController())
	}
}
