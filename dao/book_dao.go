package dao

import "database/sql"

func NewBookDao(
	db *sql.DB,
) BookDao {
	return BookDao{
		db: db,
	}
}

type BookDao struct {
	db *sql.DB
}
