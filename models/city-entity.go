package models

import (
	"gorm.io/gorm"
)

type EntityCities struct {
	gorm.Model
	Name string `gorm:"type:varchar(30)"`
	Users *[]EntityUsers `gorm:"many2many:user_cities;"`
	Books []EntityBooks `gorm:"foreignKey:CityID"`
}