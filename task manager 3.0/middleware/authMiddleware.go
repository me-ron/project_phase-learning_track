package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(c *gin.Context) {

	// check for the authorization
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"msg": "NOT AUTHORIZED"})
		c.Abort()
		return
	}

	// split the auth
	authSplit := strings.Split(auth, " ")

	if len(authSplit) != 2 || strings.ToLower(authSplit[0]) != "bearer" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"msg": "NOT AUTHORIZED"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(authSplit[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid JWT"})
		c.Abort()
		return
	}

	// move to the next step
	c.Next()
}
