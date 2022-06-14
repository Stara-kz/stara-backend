package createCityHandler

import (
	"net/http"

	createCityController "github.com/KadirbekSharau/bookswap-backend/controllers/city/create-city"
	"github.com/KadirbekSharau/bookswap-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateCityHandler(ctx *gin.Context) 
}

type handler struct {
	service createCityController.Service
}

func NewCreateCityHandler(service createCityController.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateCityHandler(ctx *gin.Context) {
	var input createCityController.InputCreateCity
	ctx.ShouldBindJSON(&input)

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

	_, errCreateCity := h.service.CreateCityService(&input)

	switch errCreateCity {
	case "CREATE_CITY_CONFLICT_409":
		util.APIResponse(ctx, "Name field already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_CITY_FAILED_403":
		util.APIResponse(ctx, "Create new city failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create new city successfully", http.StatusCreated, http.MethodPost, nil)
	}
}