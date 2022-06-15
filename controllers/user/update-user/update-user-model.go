package updateUserController

type InputUpdateUserInfo struct {
	UserID uint `json:"user_id" validate:"required,uuid"`
	CityID uint `json:"city_id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	AvatarUrl string `json:"avatar_url"`
}