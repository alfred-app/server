package bidlist

type CreateBidListBody struct {
	TalentID   string `json:"talentID"`
	JobID      string `json:"jobID"`
	PriceOnBid int    `json:"priceOnBid"`
	BidPlaced  string `json:"bidPlaced"`
}

type Response struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
