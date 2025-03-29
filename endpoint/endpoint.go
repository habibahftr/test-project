package endpoint

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"test/services/book_service"
	"test/services/loan_service"
	"test/services/session"
)

func NewEndpoint(
	router *gin.Engine,
	db *sql.DB,
	bookService book_service.BooksService,
	sessionService session.SessionService,
	loanService loan_service.LoanService,
) Endpoint {
	return Endpoint{
		router:         router,
		db:             db,
		bookService:    bookService,
		sessionService: sessionService,
		loanService:    loanService,
	}

}

type Endpoint struct {
	router         *gin.Engine
	db             *sql.DB
	bookService    book_service.BooksService
	sessionService session.SessionService
	loanService    loan_service.LoanService
}

func InitEndpoints(endpoint *Endpoint) {
	endpoint.router.POST("/book", AuthMiddleware, endpoint.bookService.InsertBook)
	endpoint.router.GET("/book", AuthMiddleware, endpoint.bookService.GetListBook)
	endpoint.router.GET("/book/:id", AuthMiddleware, endpoint.bookService.GetBook)
	endpoint.router.PUT("/book/:id", AuthMiddleware, endpoint.bookService.UpdateBook)
	endpoint.router.DELETE("/book/:id", AuthMiddleware, endpoint.bookService.DeleteBook)

	endpoint.router.POST("/loan", AuthMiddleware, endpoint.loanService.LoanBook)

	endpoint.router.POST("/login", endpoint.sessionService.Login)
}
