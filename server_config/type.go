package server_config

import (
	"database/sql"
	"test/dao"
	"test/services/book_service"
	"test/services/loan_service"
	"test/services/session"
)

type serverAttribute struct {
	DBConnection *sql.DB
	ListDao      listDao
	Services     Services
}

type listDao struct {
	bookDao dao.BookDao
	userDao dao.UserDao
	loanDao dao.LoanDao
}

type Services struct {
	BookService    book_service.BooksService
	SessionService session.SessionService
	LoanService    loan_service.LoanService
}
