package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DingDong039/hms/internal/config"
	"github.com/DingDong039/hms/internal/models"
	apperrors "github.com/DingDong039/hms/pkg/errors"
)

// HospitalAPIService defines the interface for external hospital API operations
type HospitalAPIService interface {
	SearchPatient(id string) (*models.PatientSearchResponse, error)
}

// HospitalAAPIService implements HospitalAPIService for Hospital A
type HospitalAAPIService struct {
	config *config.Config
	client *http.Client
}

// NewHospitalAAPIService creates a new HospitalAAPIService
func NewHospitalAAPIService(config *config.Config) *HospitalAAPIService {
	return &HospitalAAPIService{
		config: config,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SearchPatient searches for a patient in Hospital A's API
func (s *HospitalAAPIService) SearchPatient(id string) (*models.PatientSearchResponse, error) {
	// Build the URL
	url := fmt.Sprintf("%s/patient/search/%s", s.config.HospitalAPI.HospitalABaseURL, id)

	// Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, apperrors.NewInternalServerError(err)
	}

	// Send the request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, apperrors.NewExternalAPIError(err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return nil, apperrors.NewExternalAPIError(fmt.Errorf("hospital API returned status %d", resp.StatusCode))
	}

	// Parse the response
	var patient models.PatientSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&patient); err != nil {
		return nil, apperrors.NewInternalServerError(err)
	}

	return &patient, nil
}

// MockHospitalAAPIService implements a mock version of HospitalAPIService for testing
type MockHospitalAAPIService struct{}

// NewMockHospitalAAPIService creates a new MockHospitalAAPIService
func NewMockHospitalAAPIService() *MockHospitalAAPIService {
	return &MockHospitalAAPIService{}
}

// SearchPatient returns mock patient data
func (s *MockHospitalAAPIService) SearchPatient(id string) (*models.PatientSearchResponse, error) {
	// For testing purposes, return mock data based on the ID
	if id == "1234567890123" || id == "AB1234567" {
		dob, _ := time.Parse("2006-01-02", "1990-01-01")
		return &models.PatientSearchResponse{
			FirstNameTH:  "สมชาย",
			MiddleNameTH: "",
			LastNameTH:   "ใจดี",
			FirstNameEN:  "Somchai",
			MiddleNameEN: "",
			LastNameEN:   "Jaidee",
			DateOfBirth:  dob,
			PatientHN:    "HN12345",
			NationalID:   "1234567890123",
			PassportID:   "AB1234567",
			PhoneNumber:  "0812345678",
			Email:        "somchai@example.com",
			Gender:       "M",
		}, nil
	}

	return nil, apperrors.NewNotFoundError("patient not found")
}
