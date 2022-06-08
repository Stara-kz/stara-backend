package db

import (
	"log"

	registerAuthController "github.com/KadirbekSharau/bookswap-backend/controllers/auth/register"
	"github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)

func AccountsDataMigrator(db *gorm.DB) (*models.EntityUsers) {
	registerRepository := registerAuthController.NewRepositoryRegister(db)
	registerService := registerAuthController.NewServiceRegister(registerRepository)
	admin := registerAuthController.InputUserRegister{
		Fullname: "admin1",
		Email: "admin1@gmail.com",
		Password: "admin532",
	}
	newAdmin, errAdmin := registerService.AdminRegisterService(&admin)
	if errAdmin == "REGISTER_CONFLICT_409" || errAdmin == "REGISTER_FAILED_403" {
		log.Printf(errAdmin)
	}

	return newAdmin;
}