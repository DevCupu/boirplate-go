package service

import (
	"fmt"

	"github.com/DevCupu/boirplate-go/internal/dto"
	"github.com/DevCupu/boirplate-go/internal/model"
	"github.com/DevCupu/boirplate-go/internal/repository"
	"github.com/DevCupu/boirplate-go/pkg/auth"
	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/google/uuid"
)

// UserService interface untuk user service
type UserService interface {
	CreateUser(req *dto.RegisterRequest) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateProfile(id string, req *dto.UserUpdateProfileRequest) (*model.User, error)
	ChangePassword(id string, req *dto.ChangePasswordRequest) error
	DeleteUser(id string) error
}

// userService implementasi dari UserService
type userService struct {
	repo repository.UserRepository
}

// NewUserService membuat instance baru UserService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser membuat user baru dari request
func (s *userService) CreateUser(req *dto.RegisterRequest) (*model.User, error) {
	// Validasi email belum terdaftar
	if s.repo.ExistsByEmail(req.Email) {
		logger.Warn("Email already registered: " + req.Email)
		return nil, fmt.Errorf("email already registered")
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		logger.Error("Failed to hash password")
		return nil, fmt.Errorf("failed to hash password")
	}

	user := &model.User{
		ID:       uuid.New().String(),
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
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

// UpdateProfile mengupdate profile user (name, email, phone)
func (s *userService) UpdateProfile(id string, req *dto.UserUpdateProfileRequest) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		logger.Warn("User not found: " + id)
		return nil, fmt.Errorf("user not found")
	}

	user.Name = req.Name
	user.Phone = req.Phone

	// Validasi email jika berubah
	if user.Email != req.Email {
		if s.repo.ExistsByEmail(req.Email) {
			logger.Warn("Email already registered: " + req.Email)
			return nil, fmt.Errorf("email already registered")
		}
		user.Email = req.Email
	}

	if err := s.repo.Update(user); err != nil {
		logger.Error("Failed to update profile: " + err.Error())
		return nil, fmt.Errorf("failed to update profile")
	}

	logger.Info("User profile updated successfully: " + id)
	return user, nil
}

// ChangePassword mengubah password user dengan verifikasi password lama
func (s *userService) ChangePassword(id string, req *dto.ChangePasswordRequest) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		logger.Warn("User not found: " + id)
		return fmt.Errorf("user not found")
	}

	// Verify old password
	if !auth.VerifyPassword(user.Password, req.OldPassword) {
		logger.Warn("Change password failed: invalid old password - " + id)
		return fmt.Errorf("invalid old password")
	}

	// Hash new password
	hashedPassword, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		logger.Error("Failed to hash password")
		return fmt.Errorf("failed to hash password")
	}

	// Update password
	user.Password = hashedPassword
	if err := s.repo.Update(user); err != nil {
		logger.Error("Failed to change password: " + err.Error())
		return fmt.Errorf("failed to change password")
	}

	logger.Info("User password changed successfully: " + id)
	return nil
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
