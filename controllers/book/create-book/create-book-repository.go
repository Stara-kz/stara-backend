package createBookController

import (
	"fmt"

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

	var cityUser model.EntityUserCity
	var books model.EntityBooks
	var userBooks model.EntityUserBook
	db := r.db.Model(&books)
	errorCode := make(chan string, 1)

	books.Name = input.Name
	books.Isbn = input.Isbn
	books.Authors = input.Authors

	checkBookExist := db.Debug().Select("*").Where(&books).Find(&books)

	if checkBookExist.RowsAffected > 0 {
		errorCode <- "CREATE_BOOK_CONFLICT_409"
		return &books, <-errorCode
	}

	checkCityExists := r.db.Model(&cityUser).Debug().Select("*").Where("user_id = ?", input.UserID).Find(&cityUser)

	if checkCityExists.RowsAffected < 1 {
		errorCode <- "CREATE_BOOK_CONFLICT_404"
		return &books, <-errorCode
	}


	books.Name = input.Name
	books.Authors = input.Authors
	books.ContentDescription = input.ContentDescription
	books.Isbn = input.Isbn
	books.BookConditionDescription = input.BookConditionDescription
	books.PhotoUrl = input.PhotoUrl
	books.Condition = input.Condition
	books.Availability = true
	books.CityID = cityUser.CityID
	addNewBook := db.Debug().Create(&books)

	fmt.Println(books)

	userBooks.UserID = input.UserID
	userBooks.BookID = books.ID

	checkOwnershipExists := db.Model(&userBooks).Debug().Select("*").Where(&userBooks).Find(&userBooks)
	if checkOwnershipExists.RowsAffected > 0 {
		errorCode <- "CREATE_BOOK_CONFLICT_404"
		return &books, <-errorCode
	}

	userBooks.UserID = input.UserID
	userBooks.BookID = books.ID
	fmt.Println(userBooks)

	r.db.Model(&userBooks).Debug().Create(&userBooks)

	db.Commit()

	if addNewBook.Error != nil {
		errorCode <- "CREATE_FIELD_FAILED_403"
		return &books, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &books, <-errorCode
}
