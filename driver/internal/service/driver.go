package service

type DriverRepository interface {
}

type DriverService struct {
	repo DriverRepository
}

func NewDriverService(repo DriverRepository) *DriverService {
	return &DriverService{
		repo: repo,
	}
}
