package entity

type Offer struct {
	Id       string `json:"id"`
	From     LatLng `json:"from"`
	To       LatLng `json:"to"`
	ClientId string `json:"client_id"`
	Price    Money  `json:"price"`
}

type Money struct {
	Amount   float64 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
}
