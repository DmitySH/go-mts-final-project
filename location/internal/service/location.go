package service

type LocationRepository interface {
}

type LocationService struct {
	repo LocationRepository
}

func NewLocationService(repo LocationRepository) *LocationService {
	return &LocationService{
		repo: repo,
	}
}
