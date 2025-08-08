package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DingDong039/hms/internal/handlers"
	"github.com/DingDong039/hms/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPatientService is a mock implementation of the PatientService interface
type MockPatientService struct {
	mock.Mock
}

func (m *MockPatientService) SearchPatient(ctx context.Context, req models.PatientSearchRequest, hospitalID int) (*models.PatientSearchResponse, error) {
	args := m.Called(ctx, req, hospitalID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.PatientSearchResponse), args.Error(1)
}

// MockAuthService is a mock implementation of the AuthService interface
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) CreateStaff(ctx context.Context, req models.StaffCreateRequest) (*models.Staff, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Staff), args.Error(1)
}

func (m *MockAuthService) Login(ctx context.Context, req models.StaffLoginRequest) (*models.StaffLoginResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.StaffLoginResponse), args.Error(1)
}

func (m *MockAuthService) ValidateToken(token string) (*models.JWTClaims, error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.JWTClaims), args.Error(1)
}

func TestSearchPatient_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockPatientService := new(MockPatientService)
	mockAuthService := new(MockAuthService)
	patientHandler := handlers.NewPatientHandler(mockPatientService, mockAuthService)

	// Create a test router
	router := gin.Default()
	v1 := router.Group("/api/v1")

	// Add a test middleware to set hospitalID in context
	v1.Use(func(c *gin.Context) {
		c.Set("hospitalID", 1)
		c.Next()
	})

	patientHandler.RegisterRoutes(v1)

	// Mock request data
	reqBody := models.PatientSearchRequest{
		ID: "1234567890123",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Mock service response
	mockPatient := &models.Patient{
		FirstNameTH: "สมชาย",
		LastNameTH:  "ใจดี",
		FirstNameEN: "Somchai",
		LastNameEN:  "Jaidee",
		DateOfBirth: time.Now(),
		PatientHN:   "HN12345",
		NationalID:  "1234567890123",
		PhoneNumber: "0812345678",
		Email:       "somchai@example.com",
		Gender:      "M",
		HospitalID:  1,
	}
	mockPatientService.On("SearchPatient", mock.Anything, reqBody, 1).Return(mockPatient, nil)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/patients/search", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-token") // Mock token
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response.Success)

	// Verify mock
	mockPatientService.AssertExpectations(t)
}

func TestSearchPatient_NotFound(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockPatientService := new(MockPatientService)
	mockAuthService := new(MockAuthService)
	patientHandler := handlers.NewPatientHandler(mockPatientService, mockAuthService)

	// Create a test router
	router := gin.Default()
	v1 := router.Group("/api/v1")

	// Add a test middleware to set hospitalID in context
	v1.Use(func(c *gin.Context) {
		c.Set("hospitalID", 1)
		c.Next()
	})

	patientHandler.RegisterRoutes(v1)

	// Mock request data
	reqBody := models.PatientSearchRequest{
		ID: "9999999999999", // Non-existent ID
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Mock service error
	mockPatientService.On("SearchPatient", mock.Anything, reqBody, 1).Return(nil, errors.New("patient not found"))

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/patients/search", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-token") // Mock token
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response.Success)

	// Verify mock
	mockPatientService.AssertExpectations(t)
}

func TestSearchPatient_ValidationError(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockPatientService := new(MockPatientService)
	mockAuthService := new(MockAuthService)
	patientHandler := handlers.NewPatientHandler(mockPatientService, mockAuthService)

	// Create a test router
	router := gin.Default()
	v1 := router.Group("/api/v1")

	// Add a test middleware to set hospitalID in context
	v1.Use(func(c *gin.Context) {
		c.Set("hospitalID", 1)
		c.Next()
	})

	patientHandler.RegisterRoutes(v1)

	// Invalid request (empty ID)
	reqBody := models.PatientSearchRequest{
		ID: "", // Empty ID
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/patients/search", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-token") // Mock token
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response.Success)
}
