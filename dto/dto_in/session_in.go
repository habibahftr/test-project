package dto_in

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=1"`
	Password string `json:"password" binding:"required,min=1"`
}
