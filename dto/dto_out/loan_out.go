package dto_out

import "time"

type LoanResponse struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	BookID    int64     `json:"book_id"`
	Quantity  int16     `json:"quantity"`
	DateStart string    `json:"date_start"`
	DateEnd   string    `json:"date_end"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy int64     `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}
