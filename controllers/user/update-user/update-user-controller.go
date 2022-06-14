package updateUserController

import "github.com/KadirbekSharau/bookswap-backend/models"

type Service interface {
	UpdateUserInfoService(input *InputUpdateUserInfo) (*models.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

/* Update User Info Service */
func (s *service) UpdateUserInfoService(input *InputUpdateUserInfo) (*models.EntityUsers, string) {

	result, err := s.repository.UpdateUserInfoRepository(input)

	return result, err
}