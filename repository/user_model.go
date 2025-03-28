package repository

import "database/sql"

type UserModel struct {
	ID       sql.NullInt64
	Username sql.NullString
	Password sql.NullString
	Name     sql.NullString
}
