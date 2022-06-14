package getBookController

import (
	"github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetBookByIdRepository(input *InputGetBook) (*models.EntityBooks, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetBookByIdRepository(input *InputGetBook) (*models.EntityBooks, string) {
	var book models.EntityBooks
	db := r.db.Model(&book)
	errorCode := make(chan string, 1)

	result := db.Debug().Select("*").Where("id = ?", input.ID).Find(&book)

	if result.RowsAffected < 1 {
		errorCode <- "FIELD_NOT_FOUND_404"
		return &book, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &book, <-errorCode

}

