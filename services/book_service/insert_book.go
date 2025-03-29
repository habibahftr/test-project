package book_service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/dto"
	"test/dto/dto_in"
	"test/dto/dto_out"
	"test/repository"
	"time"
)

func (s bookService) InsertBook(
	context *gin.Context,
) {
	var err error
	timeNow := time.Now()
	var dtoIn dto_in.BookRequest
	err = context.ShouldBindJSON(&dtoIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, dto.ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Failed " + err.Error(),
		})
		return
	}

	userID := context.MustGet("userID").(int64)
	bookModel := repository.BookModel{
		Name:      sql.NullString{String: dtoIn.Name},
		Quantity:  sql.NullInt16{Int16: int16(dtoIn.Quantity)},
		CreatedBy: sql.NullInt64{Int64: userID},
		CreatedAt: sql.NullTime{Time: timeNow},
		UpdatedBy: sql.NullInt64{Int64: userID},
		UpdatedAt: sql.NullTime{Time: timeNow},
	}

	var tx *sql.Tx
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		} else {
			_ = tx.Commit()
		}
	}()
	tx, err = s.db.Begin()
	if err != nil {
		return
	}

	var id int64
	id, err = s.bookDao.InsertBook(tx, bookModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.ResponseBody{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Internal Server Error",
		})
		return
	}

	result := dto_out.BookResponse{
		ID:        id,
		Name:      dtoIn.Name,
		Quantity:  dtoIn.Quantity,
		CreatedBy: bookModel.CreatedBy.Int64,
		CreatedAt: bookModel.CreatedAt.Time,
		UpdatedBy: bookModel.UpdatedBy.Int64,
		UpdatedAt: bookModel.UpdatedAt.Time,
	}

	context.JSON(http.StatusOK, dto.ResponseBody{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
	})
	return
}
