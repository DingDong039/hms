package utils

import (
	"fmt"
	"time"

	"github.com/DingDong039/hms/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims represents the claims in a JWT token
type JWTClaims struct {
	UserID     int `json:"user_id"`
	HospitalID int `json:"hospital_id"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token for a user
func GenerateToken(userID, hospitalID int, cfg config.JWTConfig) (string, int64, error) {
	// Set expiration time
	expirationTime := time.Now().Add(time.Duration(cfg.ExpireTime) * time.Hour)
	expiresAt := expirationTime.Unix()

	// Create claims
	claims := &JWTClaims{
		UserID:     userID,
		HospitalID: hospitalID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "hms-api",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresAt, nil
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string, cfg config.JWTConfig) (*JWTClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate token and extract claims
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
