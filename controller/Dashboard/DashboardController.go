package dashboardcontrollers

import (
	"net/http"
	db "mailstitch/internal/DB"
	accesstoken "mailstitch/internal/Helper/AccessToken"
	hashapi "mailstitch/internal/Helper/HashAPI"
	inouttiming "mailstitch/internal/Helper/InOutTiming"
	timeZone "mailstitch/internal/Helper/TimeZone"
	dashboardservice "mailstitch/service/Dashboard"

	"github.com/gin-gonic/gin"
)

func GetDashboardStatsController() gin.HandlerFunc {
	return func(c *gin.Context) {
		inTime := timeZone.GetPacificTime()

		idValue, idExists := c.Get("id")
		roleIdValue, roleIdExists := c.Get("roleId")

		if !idExists || !roleIdExists {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "User session not found."})
			return
		}

		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := dashboardservice.GetDashboardStats(dbConn, int(idValue.(float64)))

		payload := map[string]interface{}{
			"status":  resVal.Status,
			"message": resVal.Message,
			"data":    resVal.Data,
		}

		token := accesstoken.CreateToken(idValue, roleIdValue)
		inouttiming.InOutTiming(inTime, timeZone.GetPacificTime(), c.Request.URL.Path)

		c.JSON(resVal.StatusCode, gin.H{
			"data":  hashapi.Encrypt(payload, true, token),
			"token": token,
		})
	}
}
