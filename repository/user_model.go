package repository

import "database/sql"

type UsersModel struct {
	ID       sql.NullInt64
	Username sql.NullString
	Password sql.NullString
	Name     sql.NullString
	Email    sql.NullString
}
