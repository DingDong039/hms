package handlers

import (
	"database/sql"
	"net/http"

	"github.com/DingDong039/hms/internal/config"
	"github.com/DingDong039/hms/internal/models"
	"github.com/DingDong039/hms/internal/repositories"
	"github.com/DingDong039/hms/internal/services"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(router *gin.Engine, db *sql.DB, cfg *config.Config) {
	// Create repositories
	staffRepo := repositories.NewStaffRepository(db)
	patientRepo := repositories.NewPatientRepository(db)

	// Create services
	hospitalAPIService := services.NewMockHospitalAAPIService() // Use mock for now
	authService := services.NewAuthService(staffRepo, cfg)
	patientService := services.NewPatientService(patientRepo, hospitalAPIService)

	// Create handlers
	authHandler := NewAuthHandler(authService)
	patientHandler := NewPatientHandler(patientService, authService)

	// API version group
	v1 := router.Group("/api/v1")

	// Health check endpoint
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.NewSuccessResponse(gin.H{"status": "ok"}))
	})

	// Register routes for each handler
	authHandler.RegisterRoutes(v1)
	patientHandler.RegisterRoutes(v1)
}
