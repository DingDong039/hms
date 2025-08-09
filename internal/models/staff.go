package models

import "time"

// Staff represents a hospital staff member
type Staff struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // Password is not exposed in JSON responses
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StaffCreateRequest represents a request to create a new staff member
type StaffCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

// StaffLoginRequest represents a login request
type StaffLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// StaffLoginResponse represents a successful login response
type StaffLoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}
