package dto_in

import "time"

type BookRequest struct {
	Name      string    `json:"name" binding:"required,min=1"`
	Quantity  int       `json:"quantity" binding:"required,min=1"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteBook struct {
	UpdatedAt time.Time `json:"updated_at" binding:"required,min=1"`
}
