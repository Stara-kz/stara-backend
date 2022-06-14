package createCityController

import "github.com/KadirbekSharau/bookswap-backend/models"

type Service interface {
	CreateCityService(input *InputCreateCity) (*models.EntityCities, string)
}

type service struct {
	repository Repository
}

func NewCreateCityService(repository Repository) Service {
	return &service{repository: repository}
}

/* Create city controller */
func (s *service) CreateCityService(input *InputCreateCity) (*models.EntityCities, string) {
	city := models.EntityCities{
		Name: input.Name,
	}

	return s.repository.CreateCityRepository(&city)
}