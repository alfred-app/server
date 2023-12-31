package bidlist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBidListHandler(c *gin.Context) {
	var requestBody CreateBidListBody

	jobID := c.Param("jobID")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	response := CreateBidList(&requestBody, jobID)
	c.JSON(response.Code, response.Response)
}

func GetAllBidListHandler(c *gin.Context) {
	response := GetAllBidList()
	c.JSON(response.Code, response.Response)
}

func GetBidListByIDHandler(c *gin.Context) {
	bidListID := c.Param("bidListID")
	response := GetBidListByID(bidListID)

	c.JSON(response.Code, response.Response)
}

func DeleteBidListHandler(c *gin.Context) {
	bidListID := c.Param("bidListID")
	response := DeleteBidList(c, bidListID)
	c.JSON(response.Code, response.Response)
}

func GetBidListByJobIDHandler(c *gin.Context) {
	jobID := c.Param("jobID")
	response := GetBidListByJobID(jobID)
	c.JSON(response.Code, response.Response)
}
