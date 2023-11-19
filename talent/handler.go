package talent

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var requestBody RegisterBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	response := RegisterTalent(&requestBody)
	c.JSON(response.Code, response.Response)
}

func LoginHandler(c *gin.Context) {
	var requestBody LoginBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	response := LoginTalent(&requestBody)
	c.JSON(response.Code, response.Response)
}
