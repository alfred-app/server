package bid

import (
	"alfred/database"
	"net/http"

	"github.com/google/uuid"
)

func CreateBid(data *CreateBidBody) Response {
	var bid database.BidList
	db := database.InitDB()

	parsedTalentID, talentError := uuid.Parse(data.TalentID)
	parsedJobID, jobError := uuid.Parse(data.JobID)

	if talentError != nil || jobError != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error parsing ID"}
	}

	bid = database.BidList{
		TalentID:   parsedTalentID,
		JobID:      parsedJobID,
		PriceOnBid: data.PriceOnBid,
	}

	err := db.Create(&bid).Error
	if err != nil {
		return Response{Code: http.StatusInternalServerError, Response: "Error Create Bid"}
	}

	return Response{Code: http.StatusCreated, Response: bid}
}
