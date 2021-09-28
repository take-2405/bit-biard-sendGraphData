package dto

type ErrorResponse struct{
	ErrorMessage string `json:"errorMessage"`
}

type Response struct{
	Graph       []Graph `json:"graph"`
	Timestamp   int     `json:"timestamp"`
	NowRate     float64 `json:"nowRate"`
}

type Graph struct{
	Timestamp int     `json:"timestamp"`
	Rate      float64 `json:"rate"`
}

type GetNowJPY struct {
	Success bool   `json:"success"`
	Rate    string `json:"rate"`
	Amount  string `json:"amount"`
	Price   string `json:"price"`
}

type BitcoinRate struct{
	NowRate    float64  `json:"rate"`
	Timestamp int `json:"timeStamp"`
}

type CoinCheck struct{
	NowRate    string  `json:"rate"`
}