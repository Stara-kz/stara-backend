package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type EntityBooks struct {
	gorm.Model
	Name string `gorm:"type:varchar(30)"`
	Authors pq.StringArray `gorm:"type:varchar[];not null"`
	ContentDescription string `gorm:"type:varchar"`
	Isbn string `gorm:"type:varchar(50)"`
	PhotoUrl  string `gorm:"type:varchar"`
	BookConditionDescription string `gorm:"type:varchar"`
	Condition uint `gorm:"type:varchar"`
	Availability    bool   `gorm:"type:bool;default:true"`
	UserBooks []EntityUserBook `gorm:"foreignKey:BookID"`
	CityID uint
}