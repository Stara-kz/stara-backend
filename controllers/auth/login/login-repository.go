package loginAuthController

import (
	util "github.com/KadirbekSharau/bookswap-backend/util"
	model "github.com/KadirbekSharau/bookswap-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	UserLoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
	AdminLoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* User Login Repository Service */
func (r *repository) UserLoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = input.Password

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	} else if !users.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}

/* Admin Login Repository Service */
func (r *repository) AdminLoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var admin model.EntityUsers
	db := r.db.Model(&admin)
	errorCode := make(chan string, 1)

	admin.Email = input.Email
	admin.Password = input.Password

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&admin)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &admin, <-errorCode
	} else if !admin.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &admin, <-errorCode
	} else if !admin.IsAdmin {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &admin, <-errorCode
	}

	comparePassword := util.ComparePassword(admin.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &admin, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &admin, <-errorCode
}