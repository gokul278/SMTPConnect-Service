package accesstoken

import (
	"fmt"
	"os"
	logger "smtpconnect/internal/Helper/Logger"
	timeZone "smtpconnect/internal/Helper/TimeZone"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// CreateToken generates a JWT token for a given user ID and expiration duration.
func CreateToken(id any, roleId any) string {
	jwtKey := []byte(os.Getenv("ACCESS_TOKEN"))
	claims := jwt.MapClaims{
		"id":     id,
		"roleId": roleId,
		// "branchId": branchid,
		"exp": time.Now().Add(24 * time.Hour).In(timeZone.MustGetPacificLocation()).Unix(), // expires after 'exp' duration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error creating token:", err)
		return "Invalid Token"
	}

	return tokenString
}

func CreateTokenWithoutExpiry(programId any) string {
	jwtKey := []byte(os.Getenv("ACCESS_TOKEN"))

	claims := jwt.MapClaims{
		"programId": programId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error creating token:", err)
		return "Invalid Token"
	}

	return tokenString
}

// ValidateJWT validates the JWT token and checks if it is expired.
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("ACCESS_TOKEN")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// ✅ Check if exp exists
		if expFloat, ok := claims["exp"].(float64); ok {

			expTime := time.Unix(int64(expFloat), 0).In(timeZone.MustGetPacificLocation())

			if timeZone.GetPacificTimeToken().After(expTime) {
				return nil, fmt.Errorf("token expired at %s", expTime.Format(time.RFC3339))
			}
		}
		// ✅ If no exp → skip expiry validation

		fmt.Println("Token valid:", claims)
	}

	return token, nil
}

// JWTMiddleware protects routes by validating JWT tokens from the Authorization header.
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		log := logger.InitLogger()

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			log.Error("Missing Token")
			c.JSON(200, gin.H{"message": "Missing token", "status": false})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Validate the JWT token
		token, err := ValidateJWT(tokenString)
		if err != nil {
			if strings.Contains(err.Error(), "token expired") {
				log.Error("Token Expired")
				c.JSON(200, gin.H{"message": "Token expired", "status": false})
				c.Abort()
				return
			}
			log.Error("Invalid Token")
			c.JSON(200, gin.H{"message": "Invalid token", "status": false})
			c.Abort()
			return
		}

		// Extract the claims (user info) and set it in the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Set the claims in the context
			c.Set("id", claims["id"])
			c.Set("roleId", claims["roleId"])
			c.Set("programId", claims["programId"])
			c.Set("token", tokenString)
		}

		// Proceed to the next handler if the token is valid
		c.Next()
	}
}
