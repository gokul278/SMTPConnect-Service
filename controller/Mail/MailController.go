package mailcontrollers

import (
	"net/http"
	db "smtpconnect/internal/DB"
	accesstoken "smtpconnect/internal/Helper/AccessToken"
	hashapi "smtpconnect/internal/Helper/HashAPI"
	inouttiming "smtpconnect/internal/Helper/InOutTiming"
	timeZone "smtpconnect/internal/Helper/TimeZone"
	mailservices "smtpconnect/service/Mail"
	mailvalidate "smtpconnect/validate/Mail"

	"github.com/gin-gonic/gin"
)

func SendMailController() gin.HandlerFunc {
	return func(c *gin.Context) {
		inTime := timeZone.GetPacificTime()

		idValue, idExists := c.Get("id")
		roleIdValue, roleIdExists := c.Get("roleId")

		if !idExists || !roleIdExists {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "User session not found."})
			return
		}

		var reqVal mailvalidate.SendMailReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request format"})
			return
		}

		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := mailservices.SendMailService(dbConn, int(idValue.(float64)), reqVal)

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

func GetMailHistoryController() gin.HandlerFunc {
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

		resVal := mailservices.GetMailHistoryService(dbConn, int(idValue.(float64)))

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
