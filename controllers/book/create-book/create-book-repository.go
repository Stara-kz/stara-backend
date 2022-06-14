package createBookController

import (
	model "github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateBookRepository(input *InputCreateBook) (*model.EntityBooks, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateBookRepository(input *InputCreateBook) (*model.EntityBooks, string) {

	var city model.EntityCities
	var books model.EntityBooks
	//var users model.EntityUsers
	db := r.db.Model(&books)
	errorCode := make(chan string, 1)

	checkBookExist := db.Debug().Select("*").Where("name = ? && isbn = ? && authors = ?", input.Name, input.Isbn, input.Authors).Find(&books)

	if checkBookExist.RowsAffected > 0 {
		errorCode <- "CREATE_FIELD_CONFLICT_409"
		return &books, <-errorCode
	}

	checkCityExists := db.Debug().Preload("Users.ID").Where("id = ?", input.UserID).Find(&city)

	if checkCityExists.RowsAffected > 0 {
		errorCode <- "CREATE_FIELD_CONFLICT_409"
		return &books, <-errorCode
	}


	books.Name = input.Name
	books.Authors = input.Authors
	books.ContentDescription = input.ContentDescription
	books.Isbn = input.Isbn
	books.PhotoUrl = input.PhotoUrl
	books.BookConditionDescription = input.BookConditionDescription
	books.Condition = input.Condition
	books.Availability = true
	books.CityID = city.ID
	addNewBook := db.Debug().Create(&books)

	//checkOwnershipExists := db.Debug().Preload("Books.ID").Where("id = ?", input.UserID).Find(&users)


	db.Commit()

	if addNewBook.Error != nil {
		errorCode <- "CREATE_FIELD_FAILED_403"
		return &books, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &books, <-errorCode
}
