package service

import (
	"fmt"
	"time"

	"github.com/DevCupu/boirplate-go/internal/model"
	"github.com/DevCupu/boirplate-go/internal/repository"
	"github.com/DevCupu/boirplate-go/pkg/auth"
	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/google/uuid"
)

// UserService interface untuk user service
type UserService interface {
	CreateUser(name, email, phone, password string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(id, name, email, phone, password string) (*model.User, error)
	DeleteUser(id string) error
	Login(email, password string) (string, *model.User, error)
}

// userService implementasi dari UserService
type userService struct {
	repo repository.UserRepository
}

// NewUserService membuat instance baru UserService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser membuat user baru
func (s *userService) CreateUser(name, email, phone, password string) (*model.User, error) {
	// Validasi email belum terdaftar
	if s.repo.ExistsByEmail(email) {
		return nil, fmt.Errorf("email already registered")
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		logger.Error("Failed to hash password")
		return nil, fmt.Errorf("failed to hash password")
	}

	user := &model.User{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: hashedPassword,
		IsActive: true,
	}

	if err := s.repo.Create(user); err != nil {
		logger.Error("Failed to create user: " + err.Error())
		return nil, fmt.Errorf("failed to create user")
	}

	logger.Info("User created successfully: " + user.Email)
	return user, nil
}

// GetUserByID mendapatkan user berdasarkan ID
func (s *userService) GetUserByID(id string) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		logger.Warn("User not found: " + id)
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

// GetAllUsers mendapatkan semua user
func (s *userService) GetAllUsers() ([]model.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		logger.Error("Failed to get users: " + err.Error())
		return nil, fmt.Errorf("failed to get users")
	}
	return users, nil
}

// UpdateUser mengupdate user
func (s *userService) UpdateUser(id, name, email, phone, password string) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		logger.Warn("User not found: " + id)
		return nil, fmt.Errorf("user not found")
	}

	user.Name = name
	user.Phone = phone

	// Validasi email jika berubah
	if user.Email != email {
		if s.repo.ExistsByEmail(email) {
			return nil, fmt.Errorf("email already registered")
		}
		user.Email = email
	}

	// Update password jika diberikan
	if password != "" {
		hashedPassword, err := auth.HashPassword(password)
		if err != nil {
			logger.Error("Failed to hash password")
			return nil, fmt.Errorf("failed to hash password")
		}
		user.Password = hashedPassword
	}

	if err := s.repo.Update(user); err != nil {
		logger.Error("Failed to update user: " + err.Error())
		return nil, fmt.Errorf("failed to update user")
	}

	logger.Info("User updated successfully: " + id)
	return user, nil
}

// DeleteUser menghapus user
func (s *userService) DeleteUser(id string) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		logger.Warn("User not found: " + id)
		return fmt.Errorf("user not found")
	}

	if err := s.repo.Delete(id); err != nil {
		logger.Error("Failed to delete user: " + err.Error())
		return fmt.Errorf("failed to delete user")
	}

	logger.Info("User deleted successfully: " + id)
	return nil
}

// Login authenticate user dan generate token
func (s *userService) Login(email, password string) (string, *model.User, error) {
	// Cari user berdasarkan email
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		logger.Warn("Login failed: user not found")
		return "", nil, fmt.Errorf("invalid email or password")
	}

	// Cek apakah user aktif
	if !user.IsActive {
		logger.Warn("Login failed: user inactive")
		return "", nil, fmt.Errorf("user account is inactive")
	}

	// Verify password
	if !auth.VerifyPassword(user.Password, password) {
		logger.Warn("Login failed: invalid password")
		return "", nil, fmt.Errorf("invalid email or password")
	}

	// Generate token (24 jam)
	token, err := auth.GenerateToken(user.ID, user.Email, 24)
	if err != nil {
		logger.Error("Failed to generate token: " + err.Error())
		return "", nil, fmt.Errorf("failed to generate token")
	}

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	s.repo.Update(user)

	logger.Info("User logged in: " + user.Email)

	return token, user, nil
}
