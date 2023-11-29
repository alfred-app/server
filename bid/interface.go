package bid

type BidPrice struct {
	PriceOnBid int `json:"price_on_bid"`
}

type CreateBidBody struct {
	TalentID   string `json:"talent_id"`
	JobID      string `json:"job_id"`
	PriceOnBid int    `json:"price_on_bid"`
}

type Response struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
