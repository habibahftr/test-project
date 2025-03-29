package repository

import "database/sql"

type LoanModel struct {
	ID        sql.NullInt64
	UserID    sql.NullInt64
	BookID    sql.NullInt64
	Quantity  sql.NullInt16
	DateStart sql.NullTime
	DateEnd   sql.NullTime
	CreatedBy sql.NullInt64
	CreatedAt sql.NullTime
	UpdatedBy sql.NullInt64
	UpdatedAt sql.NullTime
	DeletedBy sql.NullInt64
	DeletedAt sql.NullTime
}
