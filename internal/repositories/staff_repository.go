package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DingDong039/hms/internal/models"
	apperrors "github.com/DingDong039/hms/pkg/errors"
)

// StaffRepository defines the interface for staff database operations
type StaffRepository interface {
	Create(ctx context.Context, staff *models.Staff) error
	FindByUsername(ctx context.Context, username string) (*models.Staff, error)
	FindByID(ctx context.Context, id int) (*models.Staff, error)
	Update(ctx context.Context, staff *models.Staff) error
	Delete(ctx context.Context, id int) error
}

// StaffRepositoryImpl implements StaffRepository
type StaffRepositoryImpl struct {
	*BaseRepositoryImpl
}

// NewStaffRepository creates a new StaffRepositoryImpl
func NewStaffRepository(db *sql.DB) *StaffRepositoryImpl {
	return &StaffRepositoryImpl{
		BaseRepositoryImpl: NewBaseRepository(db),
	}
}

// Create inserts a new staff record into the database
func (r *StaffRepositoryImpl) Create(ctx context.Context, staff *models.Staff) error {
	query := `
		INSERT INTO staff (username, password)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	err := r.DB.QueryRowContext(
		ctx,
		query,
		staff.Username,
		staff.Password,
	).Scan(&staff.ID, &staff.CreatedAt, &staff.UpdatedAt)

	if err != nil {
		// Check for duplicate key error
		if err.Error() == "pq: duplicate key value violates unique constraint" {
			return apperrors.NewDuplicateResourceError("staff member already exists")
		}
		return apperrors.NewInternalServerError(err)
	}

	return nil
}

// FindByUsername finds a staff member by username
func (r *StaffRepositoryImpl) FindByUsername(ctx context.Context, username string) (*models.Staff, error) {
	query := `
		SELECT id, username, password, created_at, updated_at
		FROM staff
		WHERE username = $1
	`

	staff := &models.Staff{}
	err := r.DB.QueryRowContext(ctx, query, username).Scan(
		&staff.ID,
		&staff.Username,
		&staff.Password,
		&staff.CreatedAt,
		&staff.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("staff member not found")
		}
		return nil, apperrors.NewInternalServerError(err)
	}

	return staff, nil
}

// FindByID finds a staff member by ID
func (r *StaffRepositoryImpl) FindByID(ctx context.Context, id int) (*models.Staff, error) {
	query := `
		SELECT id, username, password, created_at, updated_at
		FROM staff
		WHERE id = $1
	`

	staff := &models.Staff{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&staff.ID,
		&staff.Username,
		&staff.Password,
		&staff.CreatedAt,
		&staff.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("staff member not found")
		}
		return nil, apperrors.NewInternalServerError(err)
	}

	return staff, nil
}

// Update updates a staff member record
func (r *StaffRepositoryImpl) Update(ctx context.Context, staff *models.Staff) error {
	query := `
		UPDATE staff
		SET username = $1, password = $2, updated_at = $4
		WHERE id = $5
		RETURNING updated_at
	`

	now := time.Now()
	err := r.DB.QueryRowContext(
		ctx,
		query,
		staff.Username,
		staff.Password,
		now,
		staff.ID,
	).Scan(&staff.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.NewNotFoundError("staff member not found")
		}
		return apperrors.NewInternalServerError(err)
	}

	return nil
}

// Delete deletes a staff member by ID
func (r *StaffRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM staff WHERE id = $1`

	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return apperrors.NewInternalServerError(err)
	}

	if rowsAffected == 0 {
		return apperrors.NewNotFoundError("staff member not found")
	}

	return nil
}
