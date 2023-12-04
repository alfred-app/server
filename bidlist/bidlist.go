package bidlist

import (
	"alfred/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CreateBidList(data *CreateBidListBody, talentID string) Response {
	var bidList database.BidList
	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	parsedID, err := uuid.Parse(talentID)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error parsing talent ID"}
	}

	parsedBidPlaced, err := time.Parse(time.RFC3339, data.BidPlaced)
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error parsing bid placed"}
	}

	bidList = database.BidList{
		TalentID:   parsedID,
		JobID:      parsedID,
		PriceOnBid: data.PriceOnBid,
		BidPlaced:  parsedBidPlaced,
	}

	response := db.Create(&bidList)

	if response.Error != nil {
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

func DeleteBidList(bidListID string) Response {
	var bidList database.BidList
	db := database.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	response := db.Delete(&bidList, "ID=?", bidListID)
	if response.Error != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error deleting bid list"}
	}
	return Response{Code: http.StatusOK, Response: "Bid list deleted"}
}
