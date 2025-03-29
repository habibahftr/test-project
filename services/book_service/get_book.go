package book_service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	dto2 "test/dto"
	"test/dto/dto_out"
)

func (s bookService) GetBook(
	context *gin.Context,
) {
	var err error
	bookId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, dto2.ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Failed " + err.Error(),
		})
		return
	}

	response, err := s.bookDao.GetBookById(int64(bookId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto2.ResponseBody{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Internal Server Error",
		})
		return
	}

	if response.ID.Int64 == 0 {
		context.JSON(http.StatusBadRequest, dto2.ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Failed - Data Not Found",
		})
		return
	}

	result := dto_out.BookResponse{
		ID:        response.ID.Int64,
		Name:      response.Name.String,
		Quantity:  int(response.Quantity.Int16),
		CreatedBy: response.CreatedBy.Int64,
		CreatedAt: response.CreatedAt.Time,
		UpdatedBy: response.UpdatedBy.Int64,
		UpdatedAt: response.UpdatedAt.Time,
	}

	context.JSON(http.StatusOK, dto2.ResponseBody{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
	})
	return

}
