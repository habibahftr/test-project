package book_service

import (
	"database/sql"
	"test/dao"
)

func NewBookService(
	bookDao dao.BookDao,
	db *sql.DB,
) BooksService {
	return &bookService{
		bookDao: bookDao,
		db:      db,
	}

}

type bookService struct {
	bookDao dao.BookDao
	db      *sql.DB
}
