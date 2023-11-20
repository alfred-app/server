package talent

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

func GetTalentData(c *gin.Context) {
	talentID := c.Param("talentID")
	response := GetTalentByID(talentID)
	c.JSON(response.Code, response.Response)
}

func UpdateHandler(c *gin.Context) {
	var requestBody EditTalentBody
	talentID := c.Param("talentID")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	response := EditTalentData(talentID, &requestBody)
	c.JSON(response.Code, response.Response)
}

func ChangePasswordHandler(c *gin.Context) {
	var requestBody ChangePasswordBody
	talentID := c.Param("talentID")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	response := ChangePassword(talentID, &requestBody)
	c.JSON(response.Code, response.Response)
}
