package models

import (
	"gorm.io/gorm"
)

type EntityCities struct {
	gorm.Model
	Name string `gorm:"type:varchar(30)"`
	UserCities []EntityUserCity `gorm:"foreignKey:CityID"`
	Books []EntityBooks `gorm:"foreignKey:CityID"`
}