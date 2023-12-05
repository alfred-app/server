package bidlist

type CreateBidListBody struct {
	TalentID   string `json:"talentID"`
	PriceOnBid int    `json:"priceOnBid"`
}

type Response struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
