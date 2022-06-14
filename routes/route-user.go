package routes

import (
	updateUserController "github.com/KadirbekSharau/bookswap-backend/controllers/user/update-user"
	handlerUpdateUser "github.com/KadirbekSharau/bookswap-backend/handlers/user/update-user"
	"github.com/KadirbekSharau/bookswap-backend/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All User routes */
func InitUserRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		updateUserRepository updateUserController.Repository = updateUserController.NewRepository(db)
		updateUserService updateUserController.Service = updateUserController.NewService(updateUserRepository)
		updateUserHandler handlerUpdateUser.Handler = handlerUpdateUser.NewHandler(updateUserService)
	)

	groupRoute := route.Group("/api/v1/user")
	groupRoute.PUT("/info", middlewares.Auth([]int{1,2}), updateUserHandler.UpdateUserInfoHandler)
}