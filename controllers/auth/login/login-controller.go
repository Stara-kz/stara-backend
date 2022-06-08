package loginAuthController

import (
	model "github.com/KadirbekSharau/bookswap-backend/models"
)

type Service interface {
	UserLoginService(input *InputLogin) (*model.EntityUsers, string)
	AdminLoginService(input *InputLogin) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

/* User Login Service */
func (s *service) UserLoginService(input *InputLogin) (*model.EntityUsers, string) {

	user := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.UserLoginRepository(&user)

	return resultLogin, errLogin
}

/* Admin Login Service */
func (s *service) AdminLoginService(input *InputLogin) (*model.EntityUsers, string) {

	admin := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.AdminLoginRepository(&admin)

	return resultLogin, errLogin
}