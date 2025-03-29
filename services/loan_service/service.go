package loan_service

import (
	"database/sql"
	"test/dao"
)

func NewLoanService(
	loanDao dao.LoanDao,
	bookDao dao.BookDao,
	db *sql.DB,
) LoanService {
	return &loanService{
		loanDao: loanDao,
		bookDao: bookDao,
		db:      db,
	}

}

type loanService struct {
	loanDao dao.LoanDao
	bookDao dao.BookDao
	db      *sql.DB
}
