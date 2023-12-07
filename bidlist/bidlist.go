package bidlist

import (
	"alfred/database"
	"alfred/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateBidList(data *CreateBidListBody, jobID string) Response {
	var bidList database.BidList
	db := database.InitDB()

	fmt.Println(data)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	parsedTalentID, err := uuid.Parse(data.TalentID)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error parsing talent ID"}
	}

	parsedJobID, err := uuid.Parse(jobID)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error parsing job ID"}
	}

	newID, err := uuid.NewUUID()

	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error generating new ID"}
	}

	bidList = database.BidList{
		ID:         newID,
		TalentID:   parsedTalentID,
		JobID:      parsedJobID,
		PriceOnBid: data.PriceOnBid,
	}

	response := db.Create(&bidList)

	if response.Error != nil {
		fmt.Println(response.Error.Error())
		return Response{Code: http.StatusInternalServerError, Response: "Error creating bid list"}
	}
	return Response{Code: http.StatusCreated, Response: bidList}
}

func GetAllBidList() Response {
	var bidList []database.BidList

	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.Find(&bidList)
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error getting all bid list"}
	}
	return Response{Code: http.StatusOK, Response: bidList}
}

func GetBidListByID(bidListID string) Response {
	var bidList database.BidList
	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.First(&bidList, "ID=?", bidListID)
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error getting bid list"}
	}
	return Response{Code: http.StatusOK, Response: bidList}
}

func DeleteBidList(c *gin.Context, bidListID string) Response {
	var bidList database.BidList
	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.Delete(&bidList, "ID=?", bidListID)
	middleware.AuthorizationMiddleware(c, bidList.TalentID.String())
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error deleting bid list"}
	}
	return Response{Code: http.StatusOK, Response: "Bid list deleted"}
}

func GetBidListByJobID(jobID string) Response {
	var bidList []database.BidList

	db := database.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.Find(&bidList, "\"jobID\"=?", jobID)

	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error getting bid list"}
	}
	if response.RowsAffected == 0 {
		return Response{Code: http.StatusNotFound, Response: "No Bid List Found"}
	}
	return Response{Code: http.StatusOK, Response: bidList}
}
