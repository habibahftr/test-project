package repository

import "database/sql"

type BookModel struct {
	ID        sql.NullInt64
	Name      sql.NullString
	Quantity  sql.NullInt16
	CreatedBy sql.NullInt64
	CreatedAt sql.NullTime
	UpdatedBy sql.NullInt64
	UpdatedAt sql.NullTime
	DeletedBy sql.NullInt64
	DeletedAt sql.NullTime
}
