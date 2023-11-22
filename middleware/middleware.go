package middleware

import (
	"fmt"
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
		return []byte(jwtByte), nil
	})
	if err != nil {
		c.Abort()
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	fmt.Println(claims["email"])
	c.Set("token", token.Claims)
	c.Next()
}

func AuthorizationMiddleware(c *gin.Context) {
	data, isExist := c.Get("token")
	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token data not found"})
		c.Abort()
		return
	}
	tokenData, ok := data.(map[string]interface{})
	fmt.Println(data)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing token data"})
		c.Abort()
		return
	}
	email := tokenData["email"].(string)
	fmt.Println(email)
	c.Next()
}
