package session

import "github.com/gin-gonic/gin"

type SessionService interface {
	Login(
		context *gin.Context,
	)
}
