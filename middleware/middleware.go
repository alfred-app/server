package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticationMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	jwtString := os.Getenv("JWT_KEY")
	jwtByte := []byte(jwtString)

	if authHeader == "" || len(authHeader) < 7 || authHeader[:6] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or missing token"})
		c.Abort()
		return
	}
	tokenString := authHeader[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtByte, nil
	})
	if err != nil {
		return
	}
	c.Set("token", token.Claims)
	c.Next()
}
