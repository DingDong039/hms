package services

import (
	"context"

	"github.com/DingDong039/hms/internal/config"
	"github.com/DingDong039/hms/internal/models"
	"github.com/DingDong039/hms/internal/repositories"
	"github.com/DingDong039/hms/internal/utils"
	apperrors "github.com/DingDong039/hms/pkg/errors"
)

// AuthService defines the interface for authentication operations
type AuthService interface {
	CreateStaff(ctx context.Context, req models.StaffCreateRequest) (*models.Staff, error)
	Login(ctx context.Context, req models.StaffLoginRequest) (*models.StaffLoginResponse, error)
	ValidateToken(tokenString string) (*utils.JWTClaims, error)
}

// AuthServiceImpl implements AuthService
type AuthServiceImpl struct {
	staffRepo repositories.StaffRepository
	config    *config.Config
}

// NewAuthService creates a new AuthServiceImpl
func NewAuthService(staffRepo repositories.StaffRepository, config *config.Config) *AuthServiceImpl {
	return &AuthServiceImpl{
		staffRepo: staffRepo,
		config:    config,
	}
}

// CreateStaff creates a new staff member
func (s *AuthServiceImpl) CreateStaff(ctx context.Context, req models.StaffCreateRequest) (*models.Staff, error) {
	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, apperrors.NewInternalServerError(err)
	}

	// Create staff model
	staff := &models.Staff{
		Username:   req.Username,
		Password:   hashedPassword,
		HospitalID: req.HospitalID,
	}

	// Save to database
	if err := s.staffRepo.Create(ctx, staff); err != nil {
		return nil, err
	}

	// Don't return the password
	staff.Password = ""
	return staff, nil
}

// Login authenticates a staff member and returns a JWT token
func (s *AuthServiceImpl) Login(ctx context.Context, req models.StaffLoginRequest) (*models.StaffLoginResponse, error) {
	// Find staff by username and hospital ID
	staff, err := s.staffRepo.FindByUsername(ctx, req.Username, req.HospitalID)
	if err != nil {
		return nil, apperrors.NewUnauthorizedError("invalid credentials")
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, staff.Password) {
		return nil, apperrors.NewUnauthorizedError("invalid credentials")
	}

	// Generate JWT token
	token, expiresAt, err := utils.GenerateToken(staff.ID, staff.HospitalID, s.config.JWT)
	if err != nil {
		return nil, apperrors.NewInternalServerError(err)
	}

	return &models.StaffLoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// ValidateToken validates a JWT token and returns the claims
func (s *AuthServiceImpl) ValidateToken(tokenString string) (*utils.JWTClaims, error) {
	return utils.ValidateToken(tokenString, s.config.JWT)
}
