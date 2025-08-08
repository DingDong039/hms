package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DingDong039/hms/internal/models"
	apperrors "github.com/DingDong039/hms/pkg/errors"
)

// PatientRepository defines the interface for patient database operations
type PatientRepository interface {
	Create(ctx context.Context, patient *models.Patient) error
	FindByID(ctx context.Context, id int) (*models.Patient, error)
	FindByNationalID(ctx context.Context, nationalID string, hospitalID int) (*models.Patient, error)
	FindByPassportID(ctx context.Context, passportID string, hospitalID int) (*models.Patient, error)
	FindByHospital(ctx context.Context, hospitalID int) ([]models.Patient, error)
	Update(ctx context.Context, patient *models.Patient) error
	Delete(ctx context.Context, id int) error
}

// PatientRepositoryImpl implements PatientRepository
type PatientRepositoryImpl struct {
	*BaseRepositoryImpl
}

// NewPatientRepository creates a new PatientRepositoryImpl
func NewPatientRepository(db *sql.DB) *PatientRepositoryImpl {
	return &PatientRepositoryImpl{
		BaseRepositoryImpl: NewBaseRepository(db),
	}
}

// Create inserts a new patient record into the database
func (r *PatientRepositoryImpl) Create(ctx context.Context, patient *models.Patient) error {
	query := `
		INSERT INTO patients (
			national_id, passport_id, first_name_th, middle_name_th, last_name_th,
			first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn,
			phone_number, email, gender, hospital_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id, created_at, updated_at
	`

	err := r.DB.QueryRowContext(
		ctx,
		query,
		patient.NationalID,
		patient.PassportID,
		patient.FirstNameTH,
		patient.MiddleNameTH,
		patient.LastNameTH,
		patient.FirstNameEN,
		patient.MiddleNameEN,
		patient.LastNameEN,
		patient.DateOfBirth,
		patient.PatientHN,
		patient.PhoneNumber,
		patient.Email,
		patient.Gender,
		patient.HospitalID,
	).Scan(&patient.ID, &patient.CreatedAt, &patient.UpdatedAt)

	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	return nil
}

// FindByID finds a patient by ID
func (r *PatientRepositoryImpl) FindByID(ctx context.Context, id int) (*models.Patient, error) {
	query := `
		SELECT id, national_id, passport_id, first_name_th, middle_name_th, last_name_th,
			first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn,
			phone_number, email, gender, hospital_id, created_at, updated_at
		FROM patients
		WHERE id = $1
	`

	patient := &models.Patient{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&patient.ID,
		&patient.NationalID,
		&patient.PassportID,
		&patient.FirstNameTH,
		&patient.MiddleNameTH,
		&patient.LastNameTH,
		&patient.FirstNameEN,
		&patient.MiddleNameEN,
		&patient.LastNameEN,
		&patient.DateOfBirth,
		&patient.PatientHN,
		&patient.PhoneNumber,
		&patient.Email,
		&patient.Gender,
		&patient.HospitalID,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("patient not found")
		}
		return nil, apperrors.NewInternalServerError(err)
	}

	return patient, nil
}

// FindByNationalID finds a patient by national ID and hospital ID
func (r *PatientRepositoryImpl) FindByNationalID(ctx context.Context, nationalID string, hospitalID int) (*models.Patient, error) {
	query := `
		SELECT id, national_id, passport_id, first_name_th, middle_name_th, last_name_th,
			first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn,
			phone_number, email, gender, hospital_id, created_at, updated_at
		FROM patients
		WHERE national_id = $1 AND hospital_id = $2
	`

	patient := &models.Patient{}
	err := r.DB.QueryRowContext(ctx, query, nationalID, hospitalID).Scan(
		&patient.ID,
		&patient.NationalID,
		&patient.PassportID,
		&patient.FirstNameTH,
		&patient.MiddleNameTH,
		&patient.LastNameTH,
		&patient.FirstNameEN,
		&patient.MiddleNameEN,
		&patient.LastNameEN,
		&patient.DateOfBirth,
		&patient.PatientHN,
		&patient.PhoneNumber,
		&patient.Email,
		&patient.Gender,
		&patient.HospitalID,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("patient not found")
		}
		return nil, apperrors.NewInternalServerError(err)
	}

	return patient, nil
}

// FindByPassportID finds a patient by passport ID and hospital ID
func (r *PatientRepositoryImpl) FindByPassportID(ctx context.Context, passportID string, hospitalID int) (*models.Patient, error) {
	query := `
		SELECT id, national_id, passport_id, first_name_th, middle_name_th, last_name_th,
			first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn,
			phone_number, email, gender, hospital_id, created_at, updated_at
		FROM patients
		WHERE passport_id = $1 AND hospital_id = $2
	`

	patient := &models.Patient{}
	err := r.DB.QueryRowContext(ctx, query, passportID, hospitalID).Scan(
		&patient.ID,
		&patient.NationalID,
		&patient.PassportID,
		&patient.FirstNameTH,
		&patient.MiddleNameTH,
		&patient.LastNameTH,
		&patient.FirstNameEN,
		&patient.MiddleNameEN,
		&patient.LastNameEN,
		&patient.DateOfBirth,
		&patient.PatientHN,
		&patient.PhoneNumber,
		&patient.Email,
		&patient.Gender,
		&patient.HospitalID,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("patient not found")
		}
		return nil, apperrors.NewInternalServerError(err)
	}

	return patient, nil
}

// FindByHospital finds all patients in a hospital
func (r *PatientRepositoryImpl) FindByHospital(ctx context.Context, hospitalID int) ([]models.Patient, error) {
	query := `
		SELECT id, national_id, passport_id, first_name_th, middle_name_th, last_name_th,
			first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn,
			phone_number, email, gender, hospital_id, created_at, updated_at
		FROM patients
		WHERE hospital_id = $1
	`

	rows, err := r.DB.QueryContext(ctx, query, hospitalID)
	if err != nil {
		return nil, apperrors.NewInternalServerError(err)
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		err := rows.Scan(
			&patient.ID,
			&patient.NationalID,
			&patient.PassportID,
			&patient.FirstNameTH,
			&patient.MiddleNameTH,
			&patient.LastNameTH,
			&patient.FirstNameEN,
			&patient.MiddleNameEN,
			&patient.LastNameEN,
			&patient.DateOfBirth,
			&patient.PatientHN,
			&patient.PhoneNumber,
			&patient.Email,
			&patient.Gender,
			&patient.HospitalID,
			&patient.CreatedAt,
			&patient.UpdatedAt,
		)
		if err != nil {
			return nil, apperrors.NewInternalServerError(err)
		}
		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return nil, apperrors.NewInternalServerError(err)
	}

	return patients, nil
}

// Update updates a patient record
func (r *PatientRepositoryImpl) Update(ctx context.Context, patient *models.Patient) error {
	query := `
		UPDATE patients
		SET national_id = $1, passport_id = $2, first_name_th = $3, middle_name_th = $4, 
			last_name_th = $5, first_name_en = $6, middle_name_en = $7, last_name_en = $8, 
			date_of_birth = $9, patient_hn = $10, phone_number = $11, email = $12, 
			gender = $13, hospital_id = $14, updated_at = $15
		WHERE id = $16
		RETURNING updated_at
	`

	now := time.Now()
	err := r.DB.QueryRowContext(
		ctx,
		query,
		patient.NationalID,
		patient.PassportID,
		patient.FirstNameTH,
		patient.MiddleNameTH,
		patient.LastNameTH,
		patient.FirstNameEN,
		patient.MiddleNameEN,
		patient.LastNameEN,
		patient.DateOfBirth,
		patient.PatientHN,
		patient.PhoneNumber,
		patient.Email,
		patient.Gender,
		patient.HospitalID,
		now,
		patient.ID,
	).Scan(&patient.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.NewNotFoundError("patient not found")
		}
		return apperrors.NewInternalServerError(err)
	}

	return nil
}

// Delete deletes a patient by ID
func (r *PatientRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM patients WHERE id = $1`

	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	if rowsAffected == 0 {
		return apperrors.NewNotFoundError("patient not found")
	}

	return nil
}
