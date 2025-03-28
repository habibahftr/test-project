package repository

import "github.com/golang-jwt/jwt/v5"

type PayloadJWTToken struct {
	UserID int64 `json:"user_id"`
	Claims jwt.RegisteredClaims
}
