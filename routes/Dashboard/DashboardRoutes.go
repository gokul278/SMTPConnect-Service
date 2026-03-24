package dashboardroutes

import (
	dashboardcontrollers "mailstitch/controller/Dashboard"
	accesstoken "mailstitch/internal/Helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func InitDashboardRoutes(api *gin.RouterGroup) {
	route := api.Group("/dashboard")
	route.Use(accesstoken.JWTMiddleware())
	{
		route.GET("/stats", dashboardcontrollers.GetDashboardStatsController())
	}
}
