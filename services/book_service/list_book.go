package book_service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	dto2 "test/dto"
	"test/dto/dto_out"
)

func (s bookService) GetListBook(
	context *gin.Context,
) {
	var err error
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil {
		context.JSON(http.StatusBadRequest, dto2.ResponseAPI{
			Status:  http.StatusBadRequest,
			Message: "Failed " + err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
	if err != nil {
		context.JSON(http.StatusBadRequest, dto2.ResponseAPI{
			Status:  http.StatusBadRequest,
			Message: "Failed " + err.Error(),
		})
		return
	}

	response, err := s.bookDao.GetListDao(
		sql.NullInt64{Int64: int64(page)},
		sql.NullInt64{Int64: int64(limit)},
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto2.ResponseAPI{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Internal Server Error",
		})
		return
	}

	result := []dto_out.ListBook{}
	for _, book := range response {
		result = append(result, dto_out.ListBook{
			ID:        book.ID.Int64,
			Name:      book.Name.String,
			Quantity:  int(book.Quantity.Int16),
			CreatedBy: book.CreatedBy.Int64,
			UpdatedAt: book.UpdatedAt.Time,
		})
	}

	context.JSON(http.StatusOK, dto2.ResponseAPI{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
	})
	return

}
