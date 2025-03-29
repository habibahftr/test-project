package server_config

import (
	"test/services/book_service"
	"test/services/loan_service"
	"test/services/session"
)

func (s *serverAttribute) InitService() {
	result := Services{}

	result.BookService = book_service.NewBookService(
		s.ListDao.bookDao,
		s.DBConnection,
	)

	result.SessionService = session.NewSessionService(
		s.ListDao.userDao,
		s.DBConnection,
	)

	result.LoanService = loan_service.NewLoanService(
		s.ListDao.loanDao,
		s.ListDao.bookDao,
		s.DBConnection,
	)

	s.Services = result
}
