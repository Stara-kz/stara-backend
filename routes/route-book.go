package routes

import (
	createBookController "github.com/KadirbekSharau/bookswap-backend/controllers/book/create-book"
	getBookController "github.com/KadirbekSharau/bookswap-backend/controllers/book/get-book"
	getBooksController "github.com/KadirbekSharau/bookswap-backend/controllers/book/get-books"
	getBooksHandler "github.com/KadirbekSharau/bookswap-backend/handlers/book/get-books"
	getBookHandler "github.com/KadirbekSharau/bookswap-backend/handlers/book/get-book"
	createBookHandler "github.com/KadirbekSharau/bookswap-backend/handlers/book/create-book"
	"github.com/KadirbekSharau/bookswap-backend/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Book routes */
func InitBookRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		getBookRepository getBookController.Repository = getBookController.NewRepository(db)
		getBookService getBookController.Service = getBookController.NewService(getBookRepository)
		HandlerGetBook getBookHandler.Handler = getBookHandler.NewHandler(getBookService)

		getBooksRepository getBooksController.Repository = getBooksController.NewRepository(db)
		getBooksService getBooksController.Service = getBooksController.NewService(getBooksRepository)
		HandlerGetBooks getBooksHandler.Handler = getBooksHandler.NewHandler(getBooksService)

		createBookRepository createBookController.Repository = createBookController.NewRepository(db)
		createBookService createBookController.Service = createBookController.NewService(createBookRepository)
		HandlerCreateBook createBookHandler.Handler = createBookHandler.NewHandler(createBookService)
	)

	groupRoute := route.Group("/api/v1/book")
	groupRoute.GET("/:id", middlewares.Auth([]int{1,2}), HandlerGetBook.GetBookByIdHandler)
	groupRoute.GET("/", HandlerGetBooks.GetAllBooksHandler)
	groupRoute.POST("/", middlewares.Auth([]int{1,2}), HandlerCreateBook.CreateBookHandler)
}