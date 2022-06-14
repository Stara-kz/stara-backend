package getBookController

import "github.com/KadirbekSharau/bookswap-backend/models"

type Service interface {
	GetBookByIdService(input *InputGetBook) (*models.EntityBooks, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetBookByIdService(input *InputGetBook) (*models.EntityBooks, string) {

	return s.repository.GetBookByIdRepository(input)
}