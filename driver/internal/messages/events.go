package messages

import "time"

type CreatedTripEvent struct {
	Id              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	Datacontenttype string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            struct {
		TripId  string `json:"trip_id"`
		OfferId string `json:"offer_id"`
	} `json:"data"`
}
