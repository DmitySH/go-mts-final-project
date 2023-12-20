package converters

import (
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
)

func TripFromProto(trip *driver.Trip) entity.Trip {
	return entity.Trip{
		Id:       trip.GetId(),
		DriverId: trip.GetDriverId(),
		From:     latLngFromProto(trip.GetFrom()),
		To:       latLngFromProto(trip.GetTo()),
		Price: entity.Money{
			Amount:   trip.GetPrice().GetAmount(),
			Currency: trip.GetPrice().GetCurrency(),
		},
		Status: tripStatusFromProto(trip.GetStatus()),
	}
}

func latLngFromProto(latLng *driver.LatLngLiteral) entity.LatLng {
	return entity.LatLng{
		Lat: latLng.GetLat(),
		Lng: latLng.GetLng(),
	}
}

func tripStatusFromProto(status driver.TripStatus) entity.TripStatus {
	switch status {
	case driver.TripStatus_DRIVER_SEARCH:
		return entity.TripStatusDriverSearch
	case driver.TripStatus_DRIVER_FOUND:
		return entity.TripStatusDriverFound
	case driver.TripStatus_ON_POSITION:
		return entity.TripStatusOnPosition
	case driver.TripStatus_STARTED:
		return entity.TripStatusStarted
	case driver.TripStatus_ENDED:
		return entity.TripStatusEnded
	case driver.TripStatus_CANCELED:
		return entity.TripStatusCanceled
	default:
		return entity.TripStatusUnknown
	}
}
