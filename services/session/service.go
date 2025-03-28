package session

import (
	"database/sql"
	"test/dao"
)

func NewSessionService(
	userDao dao.UserDao,
	db *sql.DB,
) SessionService {
	return sessionService{
		userDao: userDao,
		db:      db,
	}

}

type sessionService struct {
	userDao dao.UserDao
	db      *sql.DB
}
