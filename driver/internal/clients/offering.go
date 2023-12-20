package clients

import (
	context "context"

	"github.com/google/uuid"

	entity "gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
)

// MockOfferingClient is a mock of OfferingClient interface.
type MockOfferingClient struct {
}

// NewMockOfferingClient creates a new mock instance.
func NewMockOfferingClient() *MockOfferingClient {
	mock := &MockOfferingClient{}
	return mock
}

// GetOfferByID mocks base method.
func (m *MockOfferingClient) GetOfferByID(ctx context.Context, offerID string) (entity.Offer, error) {
	return entity.Offer{
		Id:       offerID,
		From:     entity.LatLng{Lat: 1, Lng: 1},
		To:       entity.LatLng{Lat: 32, Lng: 32},
		ClientId: uuid.New().String(),
		Price: entity.Money{
			Amount:   2300.03,
			Currency: "RUB",
		},
	}, nil
}
