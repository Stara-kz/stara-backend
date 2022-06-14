package getCitiesController

import "github.com/KadirbekSharau/bookswap-backend/models"

type Service interface {
	GetCitiesService() (*[]models.EntityCities, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetCitiesService() (*[]models.EntityCities, string) {
	resultCities, errCities := s.repository.GetCitiesRepository()

	return resultCities, errCities
}