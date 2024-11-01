package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	Exp  int64  `json:"exp"`
	jwt.RegisteredClaims
}
