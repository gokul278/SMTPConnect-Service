package mainroutes

import (
	authenticationroutes "mailstitch/routes/Authentication"
	configurationroutes "mailstitch/routes/Configuration"
	dashboardroutes "mailstitch/routes/Dashboard"
	mailroutes "mailstitch/routes/Mail"
	userroutes "mailstitch/routes/User"

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
