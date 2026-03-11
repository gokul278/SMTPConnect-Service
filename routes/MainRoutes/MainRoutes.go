package mainroutes

import (
	authenticationroutes "smtpconnect/routes/Authentication"
	configurationroutes "smtpconnect/routes/Configuration"
	dashboardroutes "smtpconnect/routes/Dashboard"
	mailroutes "smtpconnect/routes/Mail"
	userroutes "smtpconnect/routes/User"

	"github.com/gin-gonic/gin"
)

func InitMainRoutes(router *gin.Engine) {

	api := router.Group("/api/v1")

	authenticationroutes.InitAuthenticationRoutes(api)
	userroutes.InitUserRoutes(api)
	configurationroutes.InitConfigurationRoutes(api)
	mailroutes.InitMailRoutes(api)
	dashboardroutes.InitDashboardRoutes(api)

	_ = api
}
