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
		return []byte(jwtByte), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token Expired"})
		c.Abort()
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	c.Set("token", claims)
	c.Next()
}

func AuthorizationMiddleware(c *gin.Context) {
	data, isExist := c.Get("token")
	clientID := c.Param("userID")
	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token data not found"})
		c.Abort()
		return
	}
	claims, ok := data.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing token data"})
		c.Abort()
		return
	}
	if claims["id"] != clientID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Forbidden Content"})
		c.Abort()
		return
	}
	c.Next()
}

func TalentGuard(c *gin.Context) {
	data, isExist := c.Get("token")
	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "token data not found"})
		c.Abort()
		return
	}
	claims, ok := data.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing token data"})
		c.Abort()
		return
	}
	if claims["role"] != "talent" {
		c.JSON(http.StatusForbidden, gin.H{"message": "Only Talents are allowed"})
		c.Abort()
		return
	}
	c.Next()
}

func ClientGuard(c *gin.Context) {
	data, isExist := c.Get("token")
	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "token data not found"})
		c.Abort()
		return
	}
	claims, ok := data.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error parsing token data"})
		c.Abort()
		return
	}
	if claims["role"] != "client" {
		c.JSON(http.StatusForbidden, gin.H{"message": "Only Clients are allowed"})
		c.Abort()
		return
	}
	c.Next()
}
