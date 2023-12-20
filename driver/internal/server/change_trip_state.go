package server

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
)

func (d *DriverServer) CancelTrip(context.Context, *driver.CancelTripRequest) (*driver.CancelTripResponse, error) {
	return &driver.CancelTripResponse{}, nil
}
func (d *DriverServer) AcceptTrip(context.Context, *driver.AcceptTripRequest) (*driver.AcceptTripResponse, error) {
	return &driver.AcceptTripResponse{}, nil
}
func (d *DriverServer) StartTrip(context.Context, *driver.AcceptTripRequest) (*driver.AcceptTripResponse, error) {
	return &driver.AcceptTripResponse{}, nil
}
