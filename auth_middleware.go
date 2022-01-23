package main

import (
	"namgay/jampa/controllers"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func authenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		mySigningKey := []byte(os.Getenv("JWT_SECRET"))
		h := header{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access, Sign in first"})
			return
		}

		authStrings := strings.Fields(h.Authorization)
		token, err := jwt.ParseWithClaims(authStrings[len(authStrings) - 1], &controllers.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		})

		if err!= nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access, Sign in first"})
			return
		}

		if claims, ok := token.Claims.(*controllers.MyCustomClaims); ok && token.Valid {
			c.Set("user_id", claims.UserId)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access, Sign in first"})
		}
	}
}

type header struct {
	Authorization string `header:"Authorization" binding:"required"`
}