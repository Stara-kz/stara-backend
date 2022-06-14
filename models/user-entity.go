package models

import (
	"time"

	"github.com/KadirbekSharau/bookswap-backend/util"
	"gorm.io/gorm"
)

type EntityUsers struct {
	gorm.Model
	Name string `gorm:"type:varchar(30)"`
	Surname string `gorm:"type:varchar(30)"`
	Phone string `gorm:"type:varchar(12)"`
	Email string `gorm:"type:varchar(50);unique;not null"`
	Password  string `gorm:"type:varchar(255)"`
	Address string `gorm:"type:varchar(255)"`
	Active    bool   `gorm:"type:bool;default:false"`
	AvatarUrl string `gorm:"type:varchar"`
	IsAdmin    bool   `gorm:"type:bool;default:false"`
	UserCities []EntityUserCity `gorm:"foreignKey:UserID"`
	UserBooks []EntityUserBook `gorm:"foreignKey:UserID"`
}

type EntityUserCity struct {
	gorm.Model
	UserID uint
	CityID uint
}

type EntityUserBook struct {
	gorm.Model
	UserID uint
	BookID uint
}

func (entity *EntityUsers) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}