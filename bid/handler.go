package bid

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBidHandler(c *gin.Context) {
	var requestBody BidPrice

	talentID := c.Param("talentID")
	jobID := c.Param("jobID")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	fullRequest := CreateBidBody{
		TalentID:   talentID,
		JobID:      jobID,
		PriceOnBid: requestBody.PriceOnBid,
	}
	response := CreateBid(&fullRequest)
	c.JSON(response.Code, response.Response)
}
