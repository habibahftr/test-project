package book_service

import (
	"github.com/gin-gonic/gin"
)

type BooksService interface {
	InsertBook(
		context *gin.Context,
	)

	GetListBook(
		context *gin.Context,
	)

	GetBook(
		context *gin.Context,
	)

	UpdateBook(
		context *gin.Context,
	)

	DeleteBook(
		context *gin.Context,
	)
}
