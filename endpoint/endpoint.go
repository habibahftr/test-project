package endpoint

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func NewEndpoint(
	router *gin.Engine,
	db *sql.DB,
) Endpoint {
	return Endpoint{
		router: router,
		db:     db,
	}

}

type Endpoint struct {
	router *gin.Engine
	db     *sql.DB
	//bookService books.Service
}

func InitEndpoints(endpoint *Endpoint) {
	endpoint.router.GET("/book/:id", AuthMiddleware)
	//router.POST("/login/create", controllers.CreateLoan(db))
}
