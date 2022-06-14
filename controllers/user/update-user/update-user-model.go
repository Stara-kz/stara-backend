package updateUserController

type InputUpdateUserInfo struct {
	UserID uint `json:"user_id" validate:"required,uuid"`
	CityID uint `json:"city_id" validate:"required,uuid"`
	Name string
	Surname string
	Phone string
	Address string
	AvatarUrl string
}