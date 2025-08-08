package errors

import (
	"errors"
	"fmt"
	"net/http"
)

// Custom error types
var (
	ErrNotFound          = errors.New("resource not found")
	ErrInvalidInput      = errors.New("invalid input")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrInternalServer    = errors.New("internal server error")
	ErrDuplicateResource = errors.New("resource already exists")
	ErrExternalAPI       = errors.New("external API error")
)

// AppError represents an application error with HTTP status code
type AppError struct {
	Err        error
	StatusCode int
	Message    string
}

// Error returns the error message
func (e *AppError) Error() string {
	return e.Message
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new AppError
func NewAppError(err error, statusCode int, message string) *AppError {
	return &AppError{
		Err:        err,
		StatusCode: statusCode,
		Message:    message,
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Err:        ErrNotFound,
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

// NewInvalidInputError creates a new invalid input error
func NewInvalidInputError(message string) *AppError {
	return &AppError{
		Err:        ErrInvalidInput,
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

// NewUnauthorizedError creates a new unauthorized error
func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Err:        ErrUnauthorized,
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	}
}

// NewForbiddenError creates a new forbidden error
func NewForbiddenError(message string) *AppError {
	return &AppError{
		Err:        ErrForbidden,
		StatusCode: http.StatusForbidden,
		Message:    message,
	}
}

// NewInternalServerError creates a new internal server error
func NewInternalServerError(err error) *AppError {
	return &AppError{
		Err:        ErrInternalServer,
		StatusCode: http.StatusInternalServerError,
		Message:    fmt.Sprintf("internal server error: %v", err),
	}
}

// NewDuplicateResourceError creates a new duplicate resource error
func NewDuplicateResourceError(message string) *AppError {
	return &AppError{
		Err:        ErrDuplicateResource,
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}

// NewExternalAPIError creates a new external API error
func NewExternalAPIError(err error) *AppError {
	return &AppError{
		Err:        ErrExternalAPI,
		StatusCode: http.StatusBadGateway,
		Message:    fmt.Sprintf("external API error: %v", err),
	}
}
