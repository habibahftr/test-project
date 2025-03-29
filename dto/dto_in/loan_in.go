package dto_in

import "time"

type LoanRequest struct {
	BookID    int64     `json:"book_id" binding:"required,min=1"`
	Quantity  int16     `json:"quantity" binding:"required,min=1"`
	DateStart time.Time `json:"date_start" binding:"required,min=1"`
	DateEnd   time.Time `json:"date_end" binding:"required,min=1"`
}
