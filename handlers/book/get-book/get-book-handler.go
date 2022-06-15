package getBookHandler

import (
	"net/http"

	getBookController "github.com/KadirbekSharau/bookswap-backend/controllers/book/get-book"
	"github.com/KadirbekSharau/bookswap-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetBookByIdHandler(ctx *gin.Context)
}

type handler struct {
	service getBookController.Service
}

func NewHandler(service getBookController.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetBookByIdHandler(ctx *gin.Context) {

	var input getBookController.InputGetBook
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

	result, errResult := h.service.GetBookByIdService(&input)

	switch errResult {

	case "RESULT_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Field data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Result Field data successfully", http.StatusOK, http.MethodGet, result)
	}
}