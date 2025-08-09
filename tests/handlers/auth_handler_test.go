package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DingDong039/hms/internal/handlers"
	"github.com/DingDong039/hms/internal/models"
	"github.com/DingDong039/hms/internal/services"
	"github.com/DingDong039/hms/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func (m *MockAuthService) ValidateToken(token string) (*utils.JWTClaims, error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*utils.JWTClaims), args.Error(1)
}

func TestCreateStaff_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockAuthService := new(MockAuthService)
	authHandler := handlers.NewAuthHandler(mockAuthService)

	// Create a test router
	router := gin.Default()
	v1 := router.Group("/api/v1")
	authHandler.RegisterRoutes(v1)

	// Mock request data
	reqBody := models.StaffCreateRequest{
		Username: "testuser",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Mock service response
	mockStaff := &models.Staff{
		ID:       1,
		Username: "testuser",
	}
	mockAuthService.On("CreateStaff", mock.Anything, reqBody).Return(mockStaff, nil)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/auth/staff/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response.Success)

	// Verify mock
	mockAuthService.AssertExpectations(t)
}

func TestCreateStaff_ValidationError(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockAuthService := new(MockAuthService)
	authHandler := handlers.NewAuthHandler(mockAuthService)

	// Create a test router
	router := gin.Default()
	v1 := router.Group("/api/v1")
	authHandler.RegisterRoutes(v1)

	// Invalid request (missing required fields)
	reqBody := models.StaffCreateRequest{
		Username: "testuser",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/auth/staff/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response.Success)
}

func TestLogin_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockAuthService := new(MockAuthService)
	authHandler := handlers.NewAuthHandler(mockAuthService)

	// Create a test router
	router := gin.Default()
	v1 := router.Group("/api/v1")
	authHandler.RegisterRoutes(v1)

	// Mock request data
	reqBody := models.StaffLoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Mock service response
	mockResponse := &models.StaffLoginResponse{
		Token:     "jwt-token",
		ExpiresAt: 1628432000,
	}
	mockAuthService.On("Login", mock.Anything, reqBody).Return(mockResponse, nil)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/auth/staff/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response.Success)

	// Verify mock
	mockAuthService.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockAuthService := new(MockAuthService)
	authHandler := handlers.NewAuthHandler(mockAuthService)

	// Create a test router
	router := gin.Default()
	v1 := router.Group("/api/v1")
	authHandler.RegisterRoutes(v1)

	// Mock request data
	reqBody := models.StaffLoginRequest{
		Username: "testuser",
		Password: "wrongpassword",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Mock service error
	mockAuthService.On("Login", mock.Anything, reqBody).Return(nil, services.NewValidationError("invalid credentials"))

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/auth/staff/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response.Success)

	// Verify mock
	mockAuthService.AssertExpectations(t)
}
