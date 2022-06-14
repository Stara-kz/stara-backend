package getCityController

import "github.com/KadirbekSharau/bookswap-backend/models"

type Service interface {
	GetCityByIdService(input *InputCity) (*models.EntityCities, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCityByIdService(input *InputCity) (*models.EntityCities, string) {

	city := InputCity{
		ID: input.ID,
	}
	return s.repository.GetCityByIdRepository(&city)
}