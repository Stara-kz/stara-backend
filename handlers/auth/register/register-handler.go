package registerHandler

import (
	"net/http"

	"github.com/KadirbekSharau/bookswap-backend/controllers/auth/register"
	"github.com/KadirbekSharau/bookswap-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	ActiveUserRegisterHandler(ctx *gin.Context)
	InactiveUserRegisterHandler(ctx *gin.Context)
	AdminRegisterHandler(ctx *gin.Context)
}

type handler struct {
	service registerAuthController.Service
}

func NewHandlerRegister(service registerAuthController.Service) *handler {
	return &handler{service: service}
}

/* Active User Register Handler */
func (h *handler) ActiveUserRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputUserRegister
	ctx.ShouldBindJSON(&input)
	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.ActiveUserRegisterService(&input)
	ErrUserRegisterHandler(resultRegister, ctx, errRegister)
}

/* Inactive User Register Handler */
func (h *handler) InactiveUserRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputUserRegister
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.InactiveUserRegisterService(&input)

	ErrUserRegisterHandler(resultRegister, ctx, errRegister)
}

/* Admin Register Handler */
func (h *handler) AdminRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputUserRegister
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.AdminRegisterService(&input)

	ErrUserRegisterHandler(resultRegister, ctx, errRegister)
}