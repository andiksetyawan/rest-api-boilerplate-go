package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		splitToken := strings.Split(bearerToken, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "not authorized"})
			c.Abort()
			return
		}

		tokenString := splitToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
		})

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("guid", claims["guid"])
			c.Set("email", claims["email"])
			c.Set("group", claims["group"])
			c.Next()
		} else {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
			c.Abort()
			return
		}
	}
}
