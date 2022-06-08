package models

import (
	"time"

	"github.com/KadirbekSharau/bookswap-backend/util"
	"gorm.io/gorm"
)

type EntityUsers struct {
	gorm.Model
	Fullname string `gorm:"type:varchar(255)"`
	PhoneNumber string `gorm:"type:varchar(12)"`
	Email string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255)"`
	Active    bool   `gorm:"type:bool;default:false"`
	IsAdmin    bool   `gorm:"type:bool;default:false"`
}

func (entity *EntityUsers) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}