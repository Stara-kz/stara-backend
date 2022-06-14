package db

import (
	"log"

	registerAuthController "github.com/KadirbekSharau/bookswap-backend/controllers/auth/register"
	createCityController "github.com/KadirbekSharau/bookswap-backend/controllers/city/create-city"
	"github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)

func AccountsDataMigrator(db *gorm.DB) (*models.EntityUsers) {
	registerRepository := registerAuthController.NewRepositoryRegister(db)
	registerService := registerAuthController.NewServiceRegister(registerRepository)
	admin := registerAuthController.InputUserRegister{
		Email: "admin1@gmail.com",
		Password: "admin532",
	}
	newAdmin, errAdmin := registerService.AdminRegisterService(&admin)
	if errAdmin == "REGISTER_CONFLICT_409" || errAdmin == "REGISTER_FAILED_403" {
		log.Println(errAdmin)
	}

	return newAdmin;
}

func CityDataMigrator(cities []string, db *gorm.DB) []string {
	cityRepository := createCityController.NewCreateCityRepository(db)
	cityCreateService := createCityController.NewCreateCityService(cityRepository)

	for _,city := range cities {
		city1 := createCityController.InputCreateCity{
			Name: city,
		}
		_, errCity := cityCreateService.CreateCityService(&city1)
		if errCity == "REGISTER_CONFLICT_409" || errCity == "REGISTER_FAILED_403" {
			log.Println(errCity)
		}
	}
	
	return cities
}