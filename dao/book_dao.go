package dao

import (
	"database/sql"
	"test/repository"
)

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

func (d BookDao) InsertBook(
	tx *sql.Tx,
	model repository.BookModel,
) (
	id int64,
	err error,
) {

	query :=
		`INSERT INTO 
			books (
				name, quantity, created_at, 
				created_by, updated_at, updated_by
			) 
			VALUES (
			        $1, $2, $3, 
			        $4, $5, $6
			) RETURNING id `

	param := []interface{}{
		model.Name.String, model.Quantity.Int16, model.CreatedAt.Time,
		model.CreatedBy.Int64, model.UpdatedAt.Time, model.UpdatedBy.Int64,
	}

	err = tx.QueryRow(query, param...).Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	return
}

func CountOffset(page int, limit int) int {
	return (page - 1) * limit
}

func (d BookDao) GetListDao(
	page sql.NullInt64,
	limit sql.NullInt64,
) (
	result []repository.BookModel,
	err error,
) {

	query :=
		` SELECT 
			id, name, quantity, 
			created_by, updated_at
		FROM 
			books `

	query += "LIMIT $1 OFFSET $2 WHERE deleted = FALSE "
	offset := CountOffset(int(page.Int64), int(limit.Int64))
	param := []interface{}{
		limit.Int64, offset,
	}

	rows, err := d.db.Query(query, param...)
	if err != nil {
		return
	}

	if rows != nil {
		defer func() {
			errorS := rows.Close()
			if errorS != nil {
				err = errorS
				return
			}
		}()
		for rows.Next() {
			var temp repository.BookModel
			err = rows.Scan(
				&temp.ID, &temp.Name, &temp.Quantity,
				&temp.CreatedBy, &temp.UpdatedAt)
			if err != nil {
				return
			}

			result = append(result, temp)
		}
	} else {
		return
	}

	return
}

func (d BookDao) GetBookById(
	id int64,
) (
	result repository.BookModel,
	err error,
) {
	query :=
		`SELECT 
			id, name, quantity, 
			created_by, created_at, updated_by, 
			updated_at 
		FROM books 
		WHERE id = $1 AND deleted = FALSE `

	param := []interface{}{id}
	err = d.db.QueryRow(query, param...).Scan(
		&result.ID, &result.Name, &result.Quantity,
		&result.CreatedBy, &result.CreatedAt, &result.UpdatedBy,
		&result.UpdatedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}

func (d BookDao) GetBookByIdForUpdate(
	id int64,
) (
	result repository.BookModel,
	err error,
) {
	query :=
		`SELECT 
			id, updated_at, created_by, 
			created_at, quantity, name
		FROM books 
		WHERE id = $1 AND deleted = FALSE FOR UPDATE`

	param := []interface{}{id}
	err = d.db.QueryRow(query, param...).Scan(
		&result.ID, &result.UpdatedAt, &result.CreatedBy,
		&result.CreatedAt, &result.Quantity, &result.Name,
	)

	if err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}

func (d BookDao) UpdateBook(
	tx *sql.Tx,
	model repository.BookModel,
) (
	err error,
) {
	query :=
		`UPDATE 
			books 
		SET 
		    name = $1, quantity = $2, updated_by = $3, 
		    updated_at = $4 
		WHERE id = $5 `

	param := []interface{}{
		model.Name.String, model.Quantity.Int16, model.UpdatedBy.Int64,
		model.UpdatedAt.Time, model.ID.Int64}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(param...)

	if err != nil {
		return
	}

	return

}

func (d BookDao) DeleteBook(
	tx *sql.Tx,
	model repository.BookModel,
) (
	err error,
) {
	query :=
		`UPDATE 
			books 
		SET 
		    deleted_by = $1, deleted_at = $2, deleted = true 
		WHERE id = $3 `

	param := []interface{}{
		model.DeletedBy.Int64, model.DeletedAt.Time, model.ID.Int64}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(param...)

	if err != nil {
		return
	}

	return

}
