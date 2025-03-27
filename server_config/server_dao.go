package server_config

import "test/dao"

func (s *serverAttribute) InitDao() {
	result := listDao{}

	result.bookDao = dao.NewBookDao(
		s.DBConnection,
	)
}
