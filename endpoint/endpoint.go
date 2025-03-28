package endpoint

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"test/services/book_service"
	"test/services/session"
)

func NewEndpoint(
	router *gin.Engine,
	db *sql.DB,
	bookService book_service.BooksService,
	sessionService session.SessionService,
) Endpoint {
	return Endpoint{
		router:         router,
		db:             db,
		bookService:    bookService,
		sessionService: sessionService,
	}

}

type Endpoint struct {
	router         *gin.Engine
	db             *sql.DB
	bookService    book_service.BooksService
	sessionService session.SessionService
}

func InitEndpoints(endpoint *Endpoint) {
	//endpoint.router.GET("/book/:id", AuthMiddleware)
	endpoint.router.POST("/book", endpoint.bookService.InsertBook)
	endpoint.router.GET("/book", endpoint.bookService.GetListBook)
	endpoint.router.GET("/book/:id", endpoint.bookService.GetBook)
	endpoint.router.PUT("/book/:id", endpoint.bookService.UpdateBook)
	endpoint.router.DELETE("/book/:id", endpoint.bookService.DeleteBook)

	endpoint.router.POST("/login", endpoint.sessionService.Login)
	//router.POST("/login/create", controllers.CreateLoan(db))
}
