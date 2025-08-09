package handlers

import (
	"net/http"

	"github.com/DingDong039/hms/internal/models"
	"github.com/DingDong039/hms/internal/services"
	"github.com/DingDong039/hms/internal/utils"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	authService services.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// RegisterRoutes registers the authentication routes
func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/staff/create", h.CreateStaff)
		auth.POST("/staff/login", h.Login)
	}
}

// CreateStaff handles staff creation requests
func (h *AuthHandler) CreateStaff(c *gin.Context) {
	var req models.StaffCreateRequest

	// Validate request
	if validationErrors := utils.ValidateRequest(c, &req); validationErrors != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, "validation failed"))
		return
	}

	// Create staff
	staff, err := h.authService.CreateStaff(c.Request.Context(), req)
	if err != nil {
		// Handle specific error types
		switch err.(type) {
		case *services.ValidationError:
			c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, models.NewErrorResponse(http.StatusInternalServerError, "failed to create staff"))
		}
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, models.NewSuccessResponse(staff))
}

// Login handles staff login requests
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.StaffLoginRequest

	// Validate request
	if validationErrors := utils.ValidateRequest(c, &req); validationErrors != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, "validation failed"))
		return
	}

	// Authenticate staff
	response, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		// Handle specific error types
		switch err.(type) {
		case *services.ValidationError:
			c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, err.Error()))
		default:
			c.JSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusUnauthorized, "invalid username or password"))
		}
		return
	}

	// Return success response with JWT token
	c.JSON(http.StatusOK, models.NewSuccessResponse(response))
}
