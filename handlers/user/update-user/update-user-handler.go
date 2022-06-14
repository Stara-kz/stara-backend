package updateUserHandler

import (
	"fmt"
	"net/http"

	"github.com/KadirbekSharau/bookswap-backend/controllers/user/update-user"
	util "github.com/KadirbekSharau/bookswap-backend/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	UpdateUserInfoHandler(ctx *gin.Context)
}

type handler struct {
	service updateUserController.Service
}

func NewHandler(service updateUserController.Service) *handler {
	return &handler{service: service}
}

/* Update User info Handler */
func (h *handler) UpdateUserInfoHandler(ctx *gin.Context) {

	var input updateUserController.InputUpdateUserInfo
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
	fmt.Println(input.UserID)

	errResponse, errCount := util.GoValidator(&input, UpdateUserInfoConfig.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	res, errUpdateOrganization := h.service.UpdateUserInfoService(&input)

	ErrUpdateUserInfoHandler(res, ctx, errUpdateOrganization)
}