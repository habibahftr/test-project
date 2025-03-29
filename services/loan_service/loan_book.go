package loan_service

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

func (s loanService) LoanBook(
	context *gin.Context,
) {
	var err error
	timeNow := time.Now()
	var dtoIn dto_in.LoanRequest
	err = context.ShouldBindJSON(&dtoIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, dto.ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Failed " + err.Error(),
		})
		return
	}

	userID := context.MustGet("userID").(int64)
	loanModel := repository.LoanModel{
		UserID:    sql.NullInt64{Int64: userID},
		BookID:    sql.NullInt64{Int64: dtoIn.BookID},
		Quantity:  sql.NullInt16{Int16: dtoIn.Quantity},
		DateStart: sql.NullTime{Time: dtoIn.DateStart},
		DateEnd:   sql.NullTime{Time: dtoIn.DateEnd},
		CreatedBy: sql.NullInt64{Int64: userID},
		CreatedAt: sql.NullTime{Time: timeNow},
		UpdatedBy: sql.NullInt64{Int64: userID},
		UpdatedAt: sql.NullTime{Time: timeNow},
	}

	bookOnDb, err := s.bookDao.GetBookByIdForUpdate(loanModel.BookID.Int64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.ResponseBody{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Internal Server Error",
		})
		return
	}

	if bookOnDb.ID.Int64 == 0 {
		context.JSON(http.StatusBadRequest, dto.ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Failed - Data Book Not Found",
		})
		return
	}

	if bookOnDb.Quantity.Int16 < loanModel.Quantity.Int16 {
		context.JSON(http.StatusBadRequest, dto.ResponseBody{
			Status:  http.StatusBadRequest,
			Message: "Failed - Quantity Book Not Available",
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

	bookQty := bookOnDb.Quantity.Int16 - loanModel.Quantity.Int16
	bookModel := repository.BookModel{
		ID:        sql.NullInt64{Int64: loanModel.BookID.Int64},
		Name:      sql.NullString{String: bookOnDb.Name.String},
		Quantity:  sql.NullInt16{Int16: bookQty},
		UpdatedBy: sql.NullInt64{Int64: userID},
		UpdatedAt: sql.NullTime{Time: timeNow},
	}

	err = s.bookDao.UpdateBook(tx, bookModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.ResponseBody{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Internal Server Error",
		})
		return
	}

	var id int64
	id, err = s.loanDao.InsertLoan(tx, loanModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.ResponseBody{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Internal Server Error",
		})
		return
	}

	result := dto_out.LoanResponse{
		ID:        id,
		UserID:    userID,
		BookID:    loanModel.BookID.Int64,
		Quantity:  loanModel.Quantity.Int16,
		DateStart: loanModel.DateStart.Time.Format(time.DateOnly),
		DateEnd:   loanModel.DateEnd.Time.Format(time.DateOnly),
		CreatedBy: userID,
		CreatedAt: loanModel.CreatedAt.Time,
		UpdatedBy: userID,
		UpdatedAt: loanModel.UpdatedAt.Time,
	}

	context.JSON(http.StatusOK, dto.ResponseBody{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
	})
	return
}
