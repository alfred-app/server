package chat

import (
	"github.com/gin-gonic/gin"
)

func GetAllChatHandler(c *gin.Context) {
	userId := c.Param("userId")
	targetId := c.Param("targetId")
	response := GetAllChat(userId, targetId)
	c.JSON(response.Code, response.Response)
}
