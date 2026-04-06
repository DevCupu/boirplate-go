package service

import (
	"fmt"
	"time"

	"github.com/DevCupu/boirplate-go/internal/dto"
	"github.com/DevCupu/boirplate-go/internal/model"
	"github.com/DevCupu/boirplate-go/internal/repository"
	"github.com/DevCupu/boirplate-go/pkg/auth"
	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/google/uuid"
)

// AuthService interface untuk auth service
type AuthService interface {
	Register(req *dto.RegisterRequest) (*model.User, error)
	Login(req *dto.LoginRequest) (string, *model.User, error)
}

// authService implementasi dari AuthService
type authService struct {
	authRepo repository.AuthRepository
}

// NewAuthService membuat instance baru AuthService
func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{authRepo: authRepo}
}

// Register membuat user baru dan return user data (tanpa token)
func (s *authService) Register(req *dto.RegisterRequest) (*model.User, error) {
	// Validasi email belum terdaftar
	if s.authRepo.ExistsByEmail(req.Email) {
		logger.Warn("Registration failed: email already registered")
		return nil, fmt.Errorf("email already registered")
	}

	// Hash password dengan bcrypt
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		logger.Error("Failed to hash password: " + err.Error())
		return nil, fmt.Errorf("failed to hash password")
	}

	// Create user object
	user := &model.User{
		ID:       uuid.New().String(),
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
		IsActive: true,
	}

	// Save ke database
	if err := s.authRepo.Create(user); err != nil {
		logger.Error("Failed to create user: " + err.Error())
		return nil, fmt.Errorf("failed to register user")
	}

	logger.Info("User registered successfully: " + user.Email)
	return user, nil
}

// Login authenticate user dengan email dan password, return JWT token
func (s *authService) Login(req *dto.LoginRequest) (string, *model.User, error) {
	// Cari user berdasarkan email
	user, err := s.authRepo.GetByEmail(req.Email)
	if err != nil {
		logger.Warn("Login failed: user not found - " + req.Email)
		return "", nil, fmt.Errorf("invalid email or password")
	}

	// Cek apakah user aktif
	if !user.IsActive {
		logger.Warn("Login failed: user inactive - " + req.Email)
		return "", nil, fmt.Errorf("user account is inactive")
	}

	// Verify password dengan bcrypt
	if !auth.VerifyPassword(user.Password, req.Password) {
		logger.Warn("Login failed: invalid password - " + req.Email)
		return "", nil, fmt.Errorf("invalid email or password")
	}

	// Generate JWT token (valid 24 jam)
	token, err := auth.GenerateToken(user.ID, user.Email, 24)
	if err != nil {
		logger.Error("Failed to generate token: " + err.Error())
		return "", nil, fmt.Errorf("failed to generate token")
	}

	// Update last login timestamp
	now := time.Now()
	user.LastLogin = &now
	if err := s.authRepo.Update(user); err != nil {
		logger.Error("Failed to update last login: " + err.Error())
		// Tidak return error disini, tetap return token
	}

	logger.Info("User logged in successfully: " + user.Email)
	return token, user, nil
}
