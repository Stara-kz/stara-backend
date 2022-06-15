package getBooksHandler

import (
	"net/http"

	"github.com/KadirbekSharau/bookswap-backend/controllers/book/get-books"
	"github.com/KadirbekSharau/bookswap-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllBooksHandler(ctx *gin.Context)
}

type handler struct {
	service getBooksController.Service
}

func NewHandler(service getBooksController.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAllBooksHandler(ctx *gin.Context) {

	fields, err := h.service.GetAllBooksService()

	switch err {

	case "RESULTS_BOOKS_NOT_FOUND_404":
		util.APIResponse(ctx, "Books data is not exists", http.StatusNotFound, http.MethodGet, nil)

	default:
		util.APIResponse(ctx, "Results Books data successfully", http.StatusOK, http.MethodGet, fields)
	}
}