package model

type Property struct {
	Id      int     `json:"id"`
	Address Adrs    `json:"Address"`
	Price   float64 `json:"Price"`
}

type Adrs struct {
	Country   string `json:"country"`
	City      string `json:"city"`
	Street    string `json:"street"`
	NumOfHome string `json:"num_of_home"`
}

type Price struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency`
}
