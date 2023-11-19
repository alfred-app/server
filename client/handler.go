package client

import (
	"net/http"

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
	clientID := c.Param("clientID")
	response := GetClientByID(clientID)
	c.JSON(response.Code, response.Response)
}

func UpdateHandler(c *gin.Context) {
	var requestBody EditClientBody
	clientID := c.Param("clientID")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	response := EditClientData(clientID, &requestBody)
	c.JSON(response.Code, response.Response)
}

func ChangePasswordHandler(c *gin.Context) {
	var requestBody ChangePasswordBody
	clientID := c.Param("clientID")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	response := ChangePassword(clientID, &requestBody)
	c.JSON(response.Code, response.Response)
}
