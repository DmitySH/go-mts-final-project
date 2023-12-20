package messages

import "time"

type AcceptedTripCommand struct {
	Id              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	Datacontenttype string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            struct {
		TripId string `json:"trip_id"`
	} `json:"data"`
}

type CreatedTripCommand struct {
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

type StartedTripCommand struct {
	Id              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	Datacontenttype string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            struct {
		TripId string `json:"trip_id"`
	} `json:"data"`
}

type CancelledTripCommand struct {
	Id              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	Datacontenttype string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            struct {
		TripId string `json:"trip_id"`
	} `json:"data"`
}

type EndedTripCommand struct {
	Id              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	Datacontenttype string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            struct {
		TripId string `json:"trip_id"`
	} `json:"data"`
}
