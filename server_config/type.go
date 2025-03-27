package server_config

import (
	"database/sql"
	"test/dao"
)

type serverAttribute struct {
	DBConnection *sql.DB
	ListDao      listDao
}

type listDao struct {
	bookDao dao.BookDao
}
