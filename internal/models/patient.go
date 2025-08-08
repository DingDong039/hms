package models

import "time"

// Patient represents a patient in the system
type Patient struct {
	ID            int       `json:"id"`
	NationalID    string    `json:"national_id"`
	PassportID    string    `json:"passport_id"`
	FirstNameTH   string    `json:"first_name_th"`
	MiddleNameTH  string    `json:"middle_name_th"`
	LastNameTH    string    `json:"last_name_th"`
	FirstNameEN   string    `json:"first_name_en"`
	MiddleNameEN  string    `json:"middle_name_en"`
	LastNameEN    string    `json:"last_name_en"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	PatientHN     string    `json:"patient_hn"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `json:"email"`
	Gender        string    `json:"gender"`
	HospitalID    int       `json:"hospital_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// PatientSearchRequest represents a request to search for patients
type PatientSearchRequest struct {
	ID string `json:"id" binding:"required"` // Can be either national_id or passport_id
}

// PatientSearchResponse represents the response from the Hospital API
type PatientSearchResponse struct {
	FirstNameTH  string    `json:"first_name_th"`
	MiddleNameTH string    `json:"middle_name_th"`
	LastNameTH   string    `json:"last_name_th"`
	FirstNameEN  string    `json:"first_name_en"`
	MiddleNameEN string    `json:"middle_name_en"`
	LastNameEN   string    `json:"last_name_en"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	PatientHN    string    `json:"patient_hn"`
	NationalID   string    `json:"national_id"`
	PassportID   string    `json:"passport_id"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Gender       string    `json:"gender"`
}
