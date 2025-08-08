package handlers

import (
	"net/http"

	"github.com/DingDong039/hms/internal/middleware"
	"github.com/DingDong039/hms/internal/models"
	"github.com/DingDong039/hms/internal/services"
	"github.com/DingDong039/hms/internal/utils"
	"github.com/gin-gonic/gin"
)

// PatientHandler handles patient-related requests
type PatientHandler struct {
	patientService services.PatientService
	authService    services.AuthService
}

// NewPatientHandler creates a new PatientHandler
func NewPatientHandler(patientService services.PatientService, authService services.AuthService) *PatientHandler {
	return &PatientHandler{
		patientService: patientService,
		authService:    authService,
	}
}

// RegisterRoutes registers the patient routes
func (h *PatientHandler) RegisterRoutes(router *gin.RouterGroup) {
	// Protected routes (require authentication)
	patients := router.Group("/patients")
	patients.Use(middleware.AuthMiddleware(h.authService))
	{
		patients.POST("/search", h.SearchPatient)
	}
}

// SearchPatient handles patient search requests
func (h *PatientHandler) SearchPatient(c *gin.Context) {
	var req models.PatientSearchRequest

	// Validate request
	if validationErrors := utils.ValidateRequest(c, &req); validationErrors != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, "validation failed"))
		return
	}

	// Get hospital ID from context (set by auth middleware)
	hospitalID, exists := c.Get("hospitalID")
	if !exists {
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse(http.StatusInternalServerError, "hospital ID not found in context"))
		return
	}

	// Search for patient
	patient, err := h.patientService.SearchPatient(c.Request.Context(), req, hospitalID.(int))
	if err != nil {
		// Handle specific error types
		c.JSON(http.StatusNotFound, models.NewErrorResponse(http.StatusNotFound, "patient not found"))
		return
	}

	// Return success response
	c.JSON(http.StatusOK, models.NewSuccessResponse(patient))
}
