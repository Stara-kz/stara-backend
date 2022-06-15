package getBooksController

import "github.com/KadirbekSharau/bookswap-backend/models"

type Service interface {
	GetAllBooksService() (*[]models.EntityBooks, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAllBooksService() (*[]models.EntityBooks, string) {
	result, err := s.repository.GetAllBooksRepository()

	return result, err
}