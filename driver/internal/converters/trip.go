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

func TripFromEntity(trip entity.Trip) *driver.Trip {
	return &driver.Trip{
		Id:       trip.Id,
		DriverId: trip.DriverId,
		From:     latLngFromEntity(trip.From),
		To:       latLngFromEntity(trip.To),
		Price: &driver.Money{
			Amount:   trip.Price.Amount,
			Currency: trip.Price.Currency,
		},
		Status: tripStatusFromEntity(trip.Status),
	}
}

func latLngFromEntity(latLng entity.LatLng) *driver.LatLngLiteral {
	return &driver.LatLngLiteral{
		Lat: latLng.Lat,
		Lng: latLng.Lng,
	}
}

func tripStatusFromEntity(status entity.TripStatus) driver.TripStatus {
	switch status {
	case entity.TripStatusDriverSearch:
		return driver.TripStatus_DRIVER_SEARCH
	case entity.TripStatusDriverFound:
		return driver.TripStatus_DRIVER_FOUND
	case entity.TripStatusOnPosition:
		return driver.TripStatus_ON_POSITION
	case entity.TripStatusStarted:
		return driver.TripStatus_STARTED
	case entity.TripStatusEnded:
		return driver.TripStatus_ENDED
	case entity.TripStatusCanceled:
		return driver.TripStatus_CANCELED
	default:
		return driver.TripStatus_UNKNOWN
	}
}
