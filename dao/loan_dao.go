package dao

import (
	"database/sql"
	"test/repository"
)

func NewLoanDao(
	db *sql.DB,
) LoanDao {
	return LoanDao{
		db: db,
	}

}

type LoanDao struct {
	db *sql.DB
}

func (d LoanDao) InsertLoan(
	tx *sql.Tx,
	model repository.LoanModel,
) (
	id int64,
	err error,
) {
	query :=
		`INSERT INTO 
			loans (
			       user_id, book_id, quantity, 
			       date_start, date_end, created_by, 
			       created_at, updated_by, updated_at
			)
			VALUES (
			        $1, $2, $3, 
			        $4, $5, $6, 
			        $7, $8, $9
			) RETURNING id `

	param := []interface{}{
		model.UserID.Int64, model.BookID.Int64, model.Quantity.Int16,
		model.DateStart.Time, model.DateEnd.Time, model.CreatedBy.Int64,
		model.CreatedAt.Time, model.UpdatedBy.Int64, model.UpdatedAt.Time,
	}

	err = tx.QueryRow(query, param...).Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	return
}
