package server

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (d *DriverServer) AcceptTrip(ctx context.Context, req *driver.AcceptTripRequest) (*driver.AcceptTripResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "can't get metadata")
	}
	userIDS := md.Get(userID)
	if len(userIDS) != 1 {
		return nil, status.Error(codes.InvalidArgument, "can't get user_id from metadata")
	}

	if _, err := uuid.Parse(userIDS[0]); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id format")
	}

	if _, err := uuid.Parse(req.GetTripId()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid trip id format")
	}

	err := d.driverService.AcceptTrip(ctx, req.GetTripId(), userIDS[0])
	if err != nil {
		if errors.Is(err, service.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		loggy.Errorln(err)
		return nil, status.Errorf(codes.Internal, "can't accept trip: %w", err)
	}

	return &driver.AcceptTripResponse{}, nil
}
func (d *DriverServer) StartTrip(ctx context.Context, req *driver.StartTripRequest) (*driver.StartTripResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "can't get metadata")
	}
	userIDS := md.Get(userID)
	if len(userIDS) != 1 {
		return nil, status.Error(codes.InvalidArgument, "can't get user_id from metadata")
	}

	if _, err := uuid.Parse(userIDS[0]); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id format")
	}

	if _, err := uuid.Parse(req.GetTripId()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid trip id format")
	}

	err := d.driverService.StartTrip(ctx, req.GetTripId(), userIDS[0])
	if err != nil {
		if errors.Is(err, service.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		loggy.Errorln(err)
		return nil, status.Errorf(codes.Internal, "can't start trip: %w", err)
	}

	return &driver.StartTripResponse{}, nil
}

func (d *DriverServer) CancelTrip(ctx context.Context, req *driver.CancelTripRequest) (*driver.CancelTripResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "can't get metadata")
	}
	userIDS := md.Get(userID)
	if len(userIDS) != 1 {
		return nil, status.Error(codes.InvalidArgument, "can't get user_id from metadata")
	}

	if _, err := uuid.Parse(userIDS[0]); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id format")
	}

	if _, err := uuid.Parse(req.GetTripId()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid trip id format")
	}

	err := d.driverService.CancelTrip(ctx, req.GetTripId(), userIDS[0])
	if err != nil {
		if errors.Is(err, service.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		loggy.Errorln(err)
		return nil, status.Errorf(codes.Internal, "can't cancel trip: %w", err)
	}

	return &driver.CancelTripResponse{}, nil
}

func (d *DriverServer) EndTrip(ctx context.Context, req *driver.EndTripRequest) (*driver.EndTripResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "can't get metadata")
	}
	userIDS := md.Get(userID)
	if len(userIDS) != 1 {
		return nil, status.Error(codes.InvalidArgument, "can't get user_id from metadata")
	}

	if _, err := uuid.Parse(userIDS[0]); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id format")
	}

	if _, err := uuid.Parse(req.GetTripId()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid trip id format")
	}

	err := d.driverService.EndTrip(ctx, req.GetTripId(), userIDS[0])
	if err != nil {
		if errors.Is(err, service.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		loggy.Errorln(err)
		return nil, status.Errorf(codes.Internal, "can't end trip: %w", err)
	}

	return &driver.EndTripResponse{}, nil
}
