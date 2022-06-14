package updateUserController

import (
	"fmt"

	"github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateUserInfoRepository(input *InputUpdateUserInfo) (*models.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Update User Info Repository */
func (r *repository) UpdateUserInfoRepository(input *InputUpdateUserInfo) (*models.EntityUsers, string) {
	var users models.EntityUsers
	var userCity models.EntityUserCity
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Name = input.Name
	users.Surname = input.Surname
	users.Phone = input.Phone
	users.Address = input.Address
	users.AvatarUrl = input.AvatarUrl

	result := db.Debug().Select("name", "surname", "phone", "address", "avatar_url").Where("id = ?", input.UserID).Updates(&users).Find(&users)


	if result.RowsAffected < 1 {
		errorCode <- "RESULTS_ORGANIZATION_NOT_FOUND_404"
		return &users, <-errorCode
	}else if result.Error != nil {
		errorCode <- "RESULTS_ORGANIZATION_NOT_FOUND_404"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	checkCityExists := r.db.Model(&userCity).Debug().Select("*").Where("user_id = ?", input.UserID).Find(&userCity)

	userCity.UserID = input.UserID
	userCity.CityID = input.CityID
	fmt.Println(userCity)

	if checkCityExists.RowsAffected < 1 {
		r.db.Model(&userCity).Debug().Create(&userCity)
	} else {
		r.db.Model(&userCity).Debug().Select("user_id, city_id").Where("user_id = ?", input.UserID).Updates(&userCity)
	}

	return &users, <-errorCode
}