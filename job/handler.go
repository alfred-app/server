package job

import (
	"fmt"
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
	fmt.Println(jobID)

	c.JSON(response.Code, response.Response)
}

func GetAllJobHandler(c *gin.Context) {
	response := GetAllJobs()

	c.JSON(response.Code, response.Response)
}
