package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	mainroutes "smtpconnect/routes/MainRoutes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	// Load the DotENV
	err := godotenv.Load()
	if err != nil {
		log.Printf("❌Error loading .env file")
	}

	// ⚠️ Trust only localhost proxy (or none if you want)
	r.SetTrustedProxies(nil)

	// ✅ CORS configuration to allow only one origin
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"}, // Change to your allowed origin
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true // allow all origins dynamically
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	//Main Routes
	mainroutes.InitMainRoutes(r)

	//Ping 🎯API
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from SMPT Connect Service",
		})
	})

	//Run the Server and Log Message
	fmt.Println("✅Server is Running at Port:" + os.Getenv("PORT"))
	r.Run("0.0.0.0:" + os.Getenv("PORT"))
}
