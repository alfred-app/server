package job

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateJobHandler(c *gin.Context) {
	var requestBody CreateJobBody

	clientID := c.Param("userID")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	response := CreateJob(&requestBody, clientID)
	c.JSON(response.Code, response.Response)
}

func GetJobByIDHandler(c *gin.Context) {
	jobID := c.Param("jobID")
	response := GetJobByID(jobID)

	c.JSON(response.Code, response.Response)
}

func GetAllJobHandler(c *gin.Context) {
	response := GetAllJobs()
	c.JSON(response.Code, response.Response)
}

func GetJobByTalentIDHandler(c *gin.Context) {
	talentID := c.Param("talentID")
	response := GetJobByTalentID(talentID)
	c.JSON(response.Code, response.Response)
}

func GetJobByClientIDHandler(c *gin.Context) {
	clientID := c.Param("clientID")
	response := GetJobByClientID(clientID)
	c.JSON(response.Code, response.Response)
}

func EditJobByJobIDHandler(c *gin.Context) {
	var requestBody EditJobBody
	jobID := c.Param("jobID")
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	response := EditJobById(c, jobID, requestBody)
	c.JSON(response.Code, response.Response)
}

func SetTalentHandler(c *gin.Context) {
	var requestBody SetTalentBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	response := SetTalent(c, requestBody)
	c.JSON(response.Code, response.Response)
}
