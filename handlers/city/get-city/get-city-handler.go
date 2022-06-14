package getCityHandler

import (
	"net/http"

	getCityController "github.com/KadirbekSharau/bookswap-backend/controllers/city/get-city"
	"github.com/KadirbekSharau/bookswap-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetCityByIdHandler(ctx *gin.Context)
}

type handler struct {
	service getCityController.Service
}

func NewHandler(service getCityController.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetCityByIdHandler(ctx *gin.Context) {
	var input getCityController.InputCity

	input.ID = ctx.Params.ByName("id")

	config := util.ErrorConfig{
		Options: []util.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "ID",
				Message: "id is required on param",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	resultCity, errResultCity := h.service.GetCityByIdService(&input)

	switch errResultCity {

	case "RESULT_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "City data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Result City data successfully", http.StatusOK, http.MethodGet, resultCity)
	}
}