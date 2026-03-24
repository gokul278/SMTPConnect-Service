package authenticationcontrollers

import (
	"net/http"
	db "mailstitch/internal/DB"
	inouttiming "mailstitch/internal/Helper/InOutTiming"
	timeZone "mailstitch/internal/Helper/TimeZone"
	authenticationservices "mailstitch/service/Authentication"
	authenticationvalidate "mailstitch/validate/Authentication"

	"github.com/gin-gonic/gin"
)

func SignInController() gin.HandlerFunc {

	return func(c *gin.Context) {

		inTime := timeZone.GetPacificTime()

		var reqVal authenticationvalidate.LoginReq

		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Something went wrong, please try again, Try Again " + err.Error(),
			})
			return
		}

		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := authenticationservices.SignInService(dbConn, reqVal)

		response := gin.H{
			"status":  resVal.Status,
			"message": resVal.Message,
		}

		if resVal.Status {
			response["roleId"] = resVal.RoleId
			response["passChangeStatus"] = resVal.PassChangeStatus
			response["token"] = resVal.Token
		}

		//Server timing
		inouttiming.InOutTiming(inTime, timeZone.GetPacificTime(), c.Request.URL.Path)

		c.JSON(resVal.StatusCode, gin.H{
			"data": response,
		})

	}
}

func SignUpController() gin.HandlerFunc {

	return func(c *gin.Context) {

		inTime := timeZone.GetPacificTime()

		var reqVal authenticationvalidate.SignupReq

		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Something went wrong, please try again, Try Again " + err.Error(),
			})
			return
		}

		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := authenticationservices.SignUpService(dbConn, reqVal)

		response := gin.H{
			"status":  resVal.Status,
			"message": resVal.Message,
		}

		if resVal.Status {
			response["roleId"] = resVal.RoleId
			response["passChangeStatus"] = resVal.PassChangeStatus
			response["token"] = resVal.Token
		}

		//Server timing
		inouttiming.InOutTiming(inTime, timeZone.GetPacificTime(), c.Request.URL.Path)

		c.JSON(resVal.StatusCode, gin.H{
			"data": response,
		})

	}
}
