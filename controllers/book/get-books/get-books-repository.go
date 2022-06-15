package getBooksController

import (
	"github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)


type Repository interface {
	GetAllBooksRepository() (*[]models.EntityBooks, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllBooksRepository() (*[]models.EntityBooks, string) {
	var books []models.EntityBooks
	errorCode := make(chan string, 1)

	db := r.db.Model(&books)
	result := db.Debug().Select("*").Find(&books)

	if result.Error != nil {
		errorCode <- "RESULTS_BOOKS_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &books, <- errorCode
}