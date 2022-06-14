package updateUserHandler

import (
	"net/http"

	"github.com/KadirbekSharau/bookswap-backend/models"
	"github.com/KadirbekSharau/bookswap-backend/util"
	"github.com/gin-gonic/gin"
)


var UpdateUserInfoConfig = util.ErrorConfig{
	Options: []util.ErrorMetaConfig{
		{
			Tag:     "required",
			Field:   "ID",
			Message: "id is required on param",
		},
	},
}

/* Update user info Error Handler Function */
func ErrUpdateUserInfoHandler(res *models.EntityUsers, ctx *gin.Context, errUpdateUserInfo string) {
	switch errUpdateUserInfo {
	case "RESULTS_ORGANIZATION_NOT_FOUND_404":
		util.APIResponse(ctx, "user not found", http.StatusNotFound, http.MethodPut, nil)
		return

	default:
		util.APIResponse(ctx, "Updated user successfully", http.StatusCreated, http.MethodPut, res)
	}
}