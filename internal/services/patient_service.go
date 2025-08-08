package services

import (
	"context"
	"strings"

	"github.com/DingDong039/hms/internal/models"
	"github.com/DingDong039/hms/internal/repositories"
)

// PatientService defines the interface for patient operations
type PatientService interface {
	SearchPatient(ctx context.Context, req models.PatientSearchRequest, hospitalID int) (*models.PatientSearchResponse, error)
}

// PatientServiceImpl implements PatientService
type PatientServiceImpl struct {
	patientRepo        repositories.PatientRepository
	hospitalAPIService HospitalAPIService
}

// NewPatientService creates a new PatientServiceImpl
func NewPatientService(patientRepo repositories.PatientRepository, hospitalAPIService HospitalAPIService) *PatientServiceImpl {
	return &PatientServiceImpl{
		patientRepo:        patientRepo,
		hospitalAPIService: hospitalAPIService,
	}
}

// SearchPatient searches for a patient by ID (national ID or passport ID)
func (s *PatientServiceImpl) SearchPatient(ctx context.Context, req models.PatientSearchRequest, hospitalID int) (*models.PatientSearchResponse, error) {
	// Determine if the ID is a national ID or passport ID
	// Thai national ID is 13 digits
	// Passport IDs typically have letters
	var patient *models.Patient
	var err error

	id := strings.TrimSpace(req.ID)

	// Check if the ID is numeric and 13 digits (Thai national ID)
	if len(id) == 13 && isNumeric(id) {
		// Search by national ID
		patient, err = s.patientRepo.FindByNationalID(ctx, id, hospitalID)
	} else {
		// Search by passport ID
		patient, err = s.patientRepo.FindByPassportID(ctx, id, hospitalID)
	}

	// If patient is found in local database, return the data
	if err == nil && patient != nil {
		return &models.PatientSearchResponse{
			FirstNameTH:  patient.FirstNameTH,
			MiddleNameTH: patient.MiddleNameTH,
			LastNameTH:   patient.LastNameTH,
			FirstNameEN:  patient.FirstNameEN,
			MiddleNameEN: patient.MiddleNameEN,
			LastNameEN:   patient.LastNameEN,
			DateOfBirth:  patient.DateOfBirth,
			PatientHN:    patient.PatientHN,
			NationalID:   patient.NationalID,
			PassportID:   patient.PassportID,
			PhoneNumber:  patient.PhoneNumber,
			Email:        patient.Email,
			Gender:       patient.Gender,
		}, nil
	}

	// If patient is not found in local database, search in Hospital A API
	response, err := s.hospitalAPIService.SearchPatient(id)
	if err != nil {
		return nil, err
	}

	// Store the patient data in local database for future use
	newPatient := &models.Patient{
		NationalID:   response.NationalID,
		PassportID:   response.PassportID,
		FirstNameTH:  response.FirstNameTH,
		MiddleNameTH: response.MiddleNameTH,
		LastNameTH:   response.LastNameTH,
		FirstNameEN:  response.FirstNameEN,
		MiddleNameEN: response.MiddleNameEN,
		LastNameEN:   response.LastNameEN,
		DateOfBirth:  response.DateOfBirth,
		PatientHN:    response.PatientHN,
		PhoneNumber:  response.PhoneNumber,
		Email:        response.Email,
		Gender:       response.Gender,
		HospitalID:   hospitalID,
	}

	// Save patient to database (ignore errors as this is just caching)
	_ = s.patientRepo.Create(ctx, newPatient)

	return response, nil
}

// isNumeric checks if a string contains only digits
func isNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
