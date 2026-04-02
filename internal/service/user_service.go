package service

import (
	"fmt"

	"github.com/DevCupu/boirplate-go/internal/model"
	"github.com/DevCupu/boirplate-go/internal/repository"
	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/google/uuid"
)

// UserService interface untuk user business logic
type UserService interface {
	CreateUser(req *model.CreateUserRequest) (*model.UserResponse, error)
	GetUserByID(id string) (*model.UserResponse, error)
	GetAllUsers(limit int, offset int) ([]model.UserResponse, int64, error)
	UpdateUser(id string, req *model.UpdateUserRequest) (*model.UserResponse, error)
	DeleteUser(id string) error
}

// userService implementasi UserService
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService membuat instance UserService baru
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser membuat user baru
func (s *userService) CreateUser(req *model.CreateUserRequest) (*model.UserResponse, error) {
	// Cek apakah email sudah terdaftar
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		logger.Warn("Email already registered")
		return nil, fmt.Errorf("email already registered")
	}

	user := &model.User{
		ID:    uuid.New().String(),
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	logger.Info("User created successfully via service")
	return user.ToResponse(), nil
}

// GetUserByID mengambil user berdasarkan ID
func (s *userService) GetUserByID(id string) (*model.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

// GetAllUsers mengambil semua user dengan pagination
func (s *userService) GetAllUsers(limit int, offset int) ([]model.UserResponse, int64, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	users, total, err := s.userRepo.GetAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]model.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, *user.ToResponse())
	}

	return responses, total, nil
}

// UpdateUser mengupdate user
func (s *userService) UpdateUser(id string, req *model.UpdateUserRequest) (*model.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields jika tidak kosong
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if err := s.userRepo.Update(id, user); err != nil {
		return nil, err
	}

	logger.Info("User updated successfully via service")
	return user.ToResponse(), nil
}

// DeleteUser menghapus user
func (s *userService) DeleteUser(id string) error {
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(id)
}
