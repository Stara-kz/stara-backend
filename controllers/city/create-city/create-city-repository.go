package createCityController

import (
	"github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCityRepository(input *models.EntityCities) (*models.EntityCities, string)
}

type repository struct {
	db *gorm.DB
}

func NewCreateCityRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateCityRepository(input *models.EntityCities) (*models.EntityCities, string) {
	var city models.EntityCities
	db := r.db.Model(&city)
	errorCode := make(chan string, 1)

	checkCityExist := db.Debug().Select("*").Where("name = ?", input.Name).Find(&city)

	if checkCityExist.RowsAffected > 0 {
		errorCode <- "CREATE_CITY_CONFLICT_409"
		return &city, <-errorCode
	}

	city.Name = input.Name
	

	addNewCity := db.Debug().Create(&city)
	db.Commit()

	if addNewCity.Error != nil {
		errorCode <- "CREATE_CITY_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &city, <-errorCode
}