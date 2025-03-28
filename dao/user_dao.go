package dao

import (
	"database/sql"
	"test/repository"
)

func NewUserDao(
	db *sql.DB,
) UserDao {
	return UserDao{
		db: db,
	}
}

type UserDao struct {
	db *sql.DB
}

func (d UserDao) GetUserForLogin(
	model repository.UserModel,
) (
	result repository.UserModel,
	err error,
) {
	query :=
		`SELECT 
			id, name
		FROM users 
		WHERE username = $1 AND password = $2 
		  AND deleted = FALSE FOR UPDATE `

	param := []interface{}{
		model.Username.String, model.Password.String}
	err = d.db.QueryRow(query, param...).Scan(
		&result.ID, &result.Name,
	)

	if err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}
