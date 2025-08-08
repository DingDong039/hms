package models

// APIResponse represents a standard API response format
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

// APIError represents an error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Data:    data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code int, message string) APIResponse {
	return APIResponse{
		Success: false,
		Error: &APIError{
			Code:    code,
			Message: message,
		},
	}
}
