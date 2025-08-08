package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateRequest validates a request struct and returns validation errors
func ValidateRequest(c *gin.Context, req interface{}) []ValidationError {
	if err := c.ShouldBindJSON(req); err != nil {
		var validationErrors []ValidationError

		if verrs, ok := err.(validator.ValidationErrors); ok {
			for _, verr := range verrs {
				validationError := ValidationError{
					Field:   verr.Field(),
					Message: getValidationErrorMessage(verr),
				}
				validationErrors = append(validationErrors, validationError)
			}
		} else {
			validationErrors = append(validationErrors, ValidationError{
				Field:   "request",
				Message: "Invalid request format",
			})
		}

		return validationErrors
	}

	return nil
}

// getValidationErrorMessage returns a human-readable error message for a validation error
func getValidationErrorMessage(verr validator.FieldError) string {
	switch verr.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	default:
		return "Invalid value"
	}
}
