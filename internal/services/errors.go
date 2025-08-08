package services

// ValidationError represents an error that occurs during input validation
type ValidationError struct {
	message string
}

// NewValidationError creates a new ValidationError
func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		message: message,
	}
}

// Error returns the error message
func (e *ValidationError) Error() string {
	return e.message
}
