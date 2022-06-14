package createBookHandler

import (
	"net/http"

	createBook "github.com/KadirbekSharau/bookswap-backend/controllers/book/create-book"
	"github.com/KadirbekSharau/bookswap-backend/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateBookHandler(ctx *gin.Context)
}

type handler struct {
	service createBook.Service
}

func NewHandler(service createBook.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateBookHandler(ctx *gin.Context) {

	var input createBook.InputCreateBook
	ctx.ShouldBindJSON(&input)

	keys, ok := ctx.Get("user")

	if !ok {
		util.ValidatorErrorResponse(ctx, http.StatusUnauthorized, http.MethodPost, "user not found")
		return	
	}

	user_keys := keys.(jwt.MapClaims)
	if val, ok := user_keys["id"]; !ok {
		util.ValidatorErrorResponse(ctx, http.StatusUnauthorized, http.MethodPost, "user not found")
		return
	} else {
		input.UserID = uint(val.(float64))
	}

	config := util.ErrorConfig{
		Options: []util.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Name",
				Message: "name is required on body",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, errCreateField := h.service.CreateBookService(&input)

	switch errCreateField {
	case "CREATE_BOOK_CONFLICT_409":
		util.APIResponse(ctx, "Name Book already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_BOOK_FAILED_403":
		util.APIResponse(ctx, "Create new book account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	case "CREATE_BOOK_CONFLICT_404":
		util.APIResponse(ctx, "Book data couldn't be found", http.StatusNotFound, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create new field account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}