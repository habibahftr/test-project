package server_config

import (
	"test/services/book_service"
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
	s.Services = result
}
