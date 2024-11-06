package model

type Property struct {
	Id      int   `json:"id"`
	Address Adrs  `json:"address"`
	Price   Price `json:"price"`
	UserId  int   `json: id_user`
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
