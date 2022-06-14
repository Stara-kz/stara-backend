package registerAuthController

import (
	model "github.com/KadirbekSharau/bookswap-backend/models"
)

type Service interface {
	ActiveUserRegisterService(input *InputUserRegister) (*model.EntityUsers, string)
	AdminRegisterService(input *InputUserRegister) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

/* Active User Registration Service */
func (s *service) ActiveUserRegisterService(input *InputUserRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}
	resultRegister, errRegister := s.repository.ActiveUserRegisterRepository(&users)

	return resultRegister, errRegister
}

/* Admin Registration Service */
func (s *service) AdminRegisterService(input *InputUserRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Name: input.Name,
		Surname: input.Surname,
		Email:    input.Email,
		Password: input.Password,
	}

	resultRegister, errRegister := s.repository.AdminRegisterRepository(&users)

	return resultRegister, errRegister
}