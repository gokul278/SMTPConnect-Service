package usercontrollers

import (
	"net/http"
	db "smtpconnect/internal/DB"
	accesstoken "smtpconnect/internal/Helper/AccessToken"
	hashapi "smtpconnect/internal/Helper/HashAPI"
	inouttiming "smtpconnect/internal/Helper/InOutTiming"
	timeZone "smtpconnect/internal/Helper/TimeZone"
	userservices "smtpconnect/service/User"

	"github.com/gin-gonic/gin"
)

func UserProfileController() gin.HandlerFunc {

	return func(c *gin.Context) {

		inTime := timeZone.GetPacificTime()

		//Gathering the Datas From the Token
		idValue, idExists := c.Get("id")
		roleIdValue, roleIdExists := c.Get("roleId")

		if !idExists || !roleIdExists {
			// Handle error: ID is missing from context (e.g., middleware didn't set it)
			c.JSON(http.StatusUnauthorized, gin.H{ // Or StatusInternalServerError depending on why it's missing
				"status":  false,
				"message": "User ID, RoleID, Branch ID not found in request context.",
			})
			return
		}

		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := userservices.UserProfileService(dbConn, int(idValue.(float64)))

		payload := map[string]interface{}{
			"status":  resVal.Status,
			"message": resVal.Message,
			"data":    resVal.Data,
		}

		//Create a tokens
		token := accesstoken.CreateToken(idValue, roleIdValue)

		//Server timing
		inouttiming.InOutTiming(inTime, timeZone.GetPacificTime(), c.Request.URL.Path)

		//Send a Reponse
		c.JSON(resVal.StatusCode, gin.H{
			"data":  hashapi.Encrypt(payload, true, token),
			"token": token,
		})

	}
}
