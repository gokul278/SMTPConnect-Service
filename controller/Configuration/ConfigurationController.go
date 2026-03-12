package configurationcontrollers

import (
	"net/http"
	db "smtpconnect/internal/DB"
	accesstoken "smtpconnect/internal/Helper/AccessToken"
	hashapi "smtpconnect/internal/Helper/HashAPI"
	inouttiming "smtpconnect/internal/Helper/InOutTiming"
	timeZone "smtpconnect/internal/Helper/TimeZone"
	configurationservices "smtpconnect/service/Configuration"
	configurationvalidate "smtpconnect/validate/Configuration"

	"github.com/gin-gonic/gin"
)

func AddConfigurationController() gin.HandlerFunc {
	return func(c *gin.Context) {
		inTime := timeZone.GetPacificTime()

		// 1. Get User ID from Token Context
		idValue, idExists := c.Get("id")
		roleIdValue, roleIdExists := c.Get("roleId")

		if !idExists || !roleIdExists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "User session not found.",
			})
			return
		}

		// 2. Bind JSON Request
		var reqVal configurationvalidate.ConfigReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request format: " + err.Error(),
			})
			return
		}

		// 3. Database Connection
		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		// 4. Call Service
		resVal := configurationservices.CreateConfigurationService(dbConn, int(idValue.(float64)), reqVal)

		// 5. Prepare Response
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

func GetConfigurationsController() gin.HandlerFunc {
	return func(c *gin.Context) {
		inTime := timeZone.GetPacificTime()

		// 1. Get User ID from Token Context
		idValue, idExists := c.Get("id")
		roleIdValue, roleIdExists := c.Get("roleId")

		if !idExists || !roleIdExists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "User session not found.",
			})
			return
		}

		// 2. Database Connection
		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		// 3. Call Service
		resVal := configurationservices.GetAllConfigurationsService(dbConn, int(idValue.(float64)))

		// 4. Prepare Response
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

func UpdateConfigurationController() gin.HandlerFunc {
	return func(c *gin.Context) {
		inTime := timeZone.GetPacificTime()

		// 1. Get User ID from Token Context
		idValue, idExists := c.Get("id")
		roleIdValue, roleIdExists := c.Get("roleId")

		if !idExists || !roleIdExists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "User session not found.",
			})
			return
		}

		// 2. Bind JSON Request
		type UpdateReq struct {
			Id int `json:"id"`
			configurationvalidate.ConfigReq
		}

		var reqVal UpdateReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request format: " + err.Error(),
			})
			return
		}

		// 3. Database Connection
		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		// 4. Call Service
		resVal := configurationservices.UpdateConfigurationService(dbConn, int(idValue.(float64)), reqVal.Id, reqVal.ConfigReq)

		// 5. Prepare Response
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

func DeleteConfigurationController() gin.HandlerFunc {
	return func(c *gin.Context) {
		inTime := timeZone.GetPacificTime()

		idValue, idExists := c.Get("id")
		roleIdValue, roleIdExists := c.Get("roleId")

		if !idExists || !roleIdExists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "User session not found.",
			})
			return
		}

		type DeleteReq struct {
			Id int `json:"id"`
		}

		var reqVal DeleteReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request format: " + err.Error(),
			})
			return
		}

		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := configurationservices.DeleteConfigurationService(dbConn, int(idValue.(float64)), reqVal.Id)

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
