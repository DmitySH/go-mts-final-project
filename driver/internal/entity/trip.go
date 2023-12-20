package entity

type TripStatus string

const (
	TripStatusUnknown      TripStatus = "UNKNOWN"
	TripStatusDriverSearch TripStatus = "DRIVER_SEARCH"
	TripStatusDriverFound  TripStatus = "DRIVER_FOUND"
	TripStatusOnPosition   TripStatus = "ON_POSITION"
	TripStatusStarted      TripStatus = "STARTED"
	TripStatusEnded        TripStatus = "ENDED"
	TripStatusCanceled     TripStatus = "CANCELED"
)

type Trip struct {
	Id       string     `bson:"id"`
	DriverId string     `bson:"driver_id"`
	From     LatLng     `bson:"from"`
	To       LatLng     `bson:"to"`
	Price    Money      `bson:"money"`
	Status   TripStatus `bson:"status"`
}
