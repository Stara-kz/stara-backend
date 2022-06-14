package db

import (
	"log"

	"github.com/KadirbekSharau/bookswap-backend/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cities = []string{"Алматы", "Нур-Султан", "Актау", "Атырау", "Жанаозен", "Павлодар", "Петропавловск", "Семей", "Туркестан", "Шымкент", "Талдыкорган", "Тараз", "Уральск", "Усть-Каменогорск"}

func NewDatabaseConnection() *gorm.DB {
	dsn := "host=localhost user=kadirbeksharau password=kadr2001 dbname=template1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(
		&models.EntityUsers{},
		&models.EntityBooks{},
		&models.EntityCities{},
	)

	AccountsDataMigrator(db)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	CityDataMigrator(cities, db)

	return db
}



func CloseDB(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	database.Close()
}