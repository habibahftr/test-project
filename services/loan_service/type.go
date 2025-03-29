package loan_service

import "github.com/gin-gonic/gin"

type LoanService interface {
	LoanBook(
		context *gin.Context,
	)
}
