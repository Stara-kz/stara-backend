package getCitiesController

import (
	"github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)


type Repository interface {
	GetCitiesRepository() (*[]models.EntityCities, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetCitiesRepository() (*[]models.EntityCities, string) {
	var cities []models.EntityCities
	errorCode := make(chan string, 1)

	db := r.db.Model(&cities)
	result := db.Debug().Select("*").Find(&cities)

	if result.Error != nil {
		errorCode <- "RESULTS_CITY_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &cities, <- errorCode
}