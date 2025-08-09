package models

import "github.com/golang-jwt/jwt/v5"

// JWTClaims represents the claims in a JWT token
type JWTClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
