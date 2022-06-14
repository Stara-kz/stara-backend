package createBookController


import (
	model "github.com/KadirbekSharau/bookswap-backend/models"
)

type Service interface {
	CreateBookService(input *InputCreateBook) (*model.EntityBooks, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateBookService(input *InputCreateBook) (*model.EntityBooks, string) {


	resultCreateField, errCreateField := s.repository.CreateBookRepository(input)
	//checkOwnershipExists := db.Debug().Preload("Books.ID").Where("id = ?", input.UserID).Find(&users)
	

	return resultCreateField, errCreateField
}