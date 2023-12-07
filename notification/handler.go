package notification

import "github.com/gin-gonic/gin"

func GetAllNotificationHandler(c *gin.Context) {
	userId := c.Param("userId")
	response := GetAllNotifications(userId)
	c.JSON(response.Code, response.Response)
}
