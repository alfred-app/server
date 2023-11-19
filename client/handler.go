package client

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var requestBody RegisterBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	response := RegisterClient(&requestBody)
	c.JSON(response.Code, response.Response)
}

func LoginHandler(c *gin.Context) {
	var requestBody LoginBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	response := LoginClient(&requestBody)
	c.JSON(response.Code, response.Response)
}

func GetClientData(c *gin.Context) {
	clientID := c.Param("id")
	response := GetClientByID(clientID)
	c.JSON(response.Code, response.Response)
}
