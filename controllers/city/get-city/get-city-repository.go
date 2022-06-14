package getCityController

import (
	"github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetCityByIdRepository(input *InputCity) (*models.EntityCities, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetCityByIdRepository(input *InputCity) (*models.EntityCities, string) {
	var city models.EntityCities
	db := r.db.Model(&city)
	errorCode := make(chan string, 1)

	resultCity := db.Debug().Select("*").Where("id = ?", input.ID).Find(&city)

	if resultCity.RowsAffected < 1 {
		errorCode <- "CITY_NOT_FOUND_404"
		return &city, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &city, <-errorCode

}

