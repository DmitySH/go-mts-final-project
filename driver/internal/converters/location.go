package converters

import (
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/location"
)

func DriverFromProto(driver *location.Driver) entity.Driver {
	return entity.Driver{
		Lat:  driver.GetLat(),
		Lng:  driver.GetLng(),
		Id:   driver.GetId(),
		Name: driver.GetName(),
		Auto: driver.GetAuto(),
	}
}
