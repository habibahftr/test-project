package book_service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	dto2 "test/dto"
	"test/dto/dto_in"
	"test/repository"
	"time"
)

func (s bookService) DeleteBook(
	context *gin.Context,
) {
	var err error
	timeNow := time.Now()
	bookId, err := strconv.Atoi(context.Param("id"))
	var dtoIn dto_in.DeleteBook
	err = context.ShouldBindJSON(&dtoIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, dto2.ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Failed " + err.Error(),
		})
		return
	}

	userID := context.MustGet("userID").(int64)
	bookModel := repository.BookModel{
		ID:        sql.NullInt64{Int64: int64(bookId)},
		DeletedBy: sql.NullInt64{Int64: userID},
		DeletedAt: sql.NullTime{Time: timeNow},
	}

	bookOnDb, err := s.bookDao.GetBookByIdForUpdate(bookModel.ID.Int64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto2.ResponseBody{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Internal Server Error",
		})
		return
	}

	if bookOnDb.ID.Int64 == 0 {
		context.JSON(http.StatusBadRequest, dto2.ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Failed - Data Not Found",
		})
		return
	}

	t1 := bookOnDb.UpdatedAt.Time.Format(time.RFC3339)
	t2 := dtoIn.UpdatedAt.Format(time.RFC3339)
	if t1 != t2 {
		context.JSON(http.StatusBadRequest, dto2.ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Failed - Data Locked",
		})
		return
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

	err = s.bookDao.DeleteBook(tx, bookModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto2.ResponseBody{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Internal Server Error",
		})
		return
	}

	context.JSON(http.StatusOK, dto2.ResponseBody{
		Status:  http.StatusOK,
		Message: "Success",
	})
}
