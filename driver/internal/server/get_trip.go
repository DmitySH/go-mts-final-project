package server

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
)

func (d *DriverServer) GetTrips(context.Context, *driver.GetTripsRequest) (*driver.GetTripsResponse, error) {
	// TODO: implement me
	return nil, nil
}
func (d *DriverServer) GetTripByID(context.Context, *driver.GetTripByIDRequest) (*driver.GetTripByIDResponse, error) {
	// TODO: implement me
	return nil, nil
}
