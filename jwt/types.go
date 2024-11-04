package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
	Exp  int64     `json:"exp"`
	jwt.RegisteredClaims
}
