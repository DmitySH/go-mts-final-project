package converters

import (
	"errors"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/pkg/api/location"
)

func LatLngLiteralFromProto(req *location.LatLngLiteral) (entity.LatLng, error) {
	if !isCorrectLat(req.GetLat()) {
		return entity.LatLng{}, errors.New("invalid lat")
	}
	if !isCorrectLng(req.GetLng()) {
		return entity.LatLng{}, errors.New("invalid lng")
	}

	return entity.LatLng{
		Lat: req.GetLat(),
		Lng: req.GetLng(),
	}, nil
}

func LatLngLiteralFromFloats(lat, lng float64) (entity.LatLng, error) {
	if !isCorrectLat(lat) {
		return entity.LatLng{}, errors.New("invalid lat")
	}
	if !isCorrectLng(lng) {
		return entity.LatLng{}, errors.New("invalid lng")
	}

	return entity.LatLng{
		Lat: lat,
		Lng: lng,
	}, nil
}

func DriversToProto(drivers []entity.Driver) []*location.Driver {
	res := make([]*location.Driver, 0, len(drivers))
	for _, driver := range drivers {
		res = append(res, &location.Driver{
			Lat:  driver.Lat,
			Lng:  driver.Lng,
			Id:   driver.Id,
			Name: driver.Name,
			Auto: driver.Auto,
		})
	}

	return res
}
