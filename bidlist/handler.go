package bidlist

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBidListHandler(c *gin.Context) {
	var requestBody CreateBidListBody

	talentID := c.Param("userID")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	response := CreateBidList(&requestBody, talentID)
	c.JSON(response.Code, response.Response)
}

func GetAllBidListHandler(c *gin.Context) {
	response := GetAllBidList()
	c.JSON(response.Code, response.Response)
}

func GetBidListByIDHandler(c *gin.Context) {
	bidListID := c.Param("bidListID")
	response := GetBidListByID(bidListID)
	fmt.Println(bidListID)

	c.JSON(response.Code, response.Response)
}

func DeleteBidListHandler(c *gin.Context) {
	bidListID := c.Param("bidListID")
	response := DeleteBidList(bidListID)
	c.JSON(response.Code, response.Response)
}
