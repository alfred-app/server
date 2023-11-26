package job

type CreateJobBody struct {
	ClientID     string  `json:"clientID"`
	Name         string  `json:"name"`
	Descriptions string  `json:"descriptions"`
	Address      string  `json:"address"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	ImageURL     string  `json:"imageURL"`
}

type Response struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
