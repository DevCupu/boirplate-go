package repository

import (
	"fmt"

	"boilerplate-go/internal/model"
	"boilerplate-go/pkg/logger"
	"gorm.io/gorm"
)

// UserRepository interface untuk user queries
type UserRepository interface {
	Create(user *model.User) error
	GetByID(id string) (*model.User, error)
	GetAll(limit int, offset int) ([]model.User, int64, error)
	Update(id string, user *model.User) error
	Delete(id string) error
	GetByEmail(email string) (*model.User, error)
}

// userRepository implementasi UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository membuat instance UserRepository baru
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create membuat user baru
func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		logger.Error("Failed to create user",)
		return fmt.Errorf("failed to create user: %w", err)
	}
	logger.Info("User created successfully", )
	return nil
}

// GetByID mengambil user berdasarkan ID
func (r *userRepository) GetByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("User not found", )
			return nil, fmt.Errorf("user not found")
		}
		logger.Error("Failed to get user by id", )
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// GetAll mengambil semua user dengan pagination
func (r *userRepository) GetAll(limit int, offset int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	if err := r.db.Model(&model.User{}).Count(&total).Error; err != nil {
		logger.Error("Failed to count users", )
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		logger.Error("Failed to get users", )
		return nil, 0, fmt.Errorf("failed to get users: %w", err)
	}

	return users, total, nil
}

// Update mengupdate user
func (r *userRepository) Update(id string, user *model.User) error {
	if err := r.db.Model(&model.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		logger.Error("Failed to update user", )
		return fmt.Errorf("failed to update user: %w", err)
	}
	logger.Info("User updated successfully", )
	return nil
}

// Delete menghapus user
func (r *userRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		logger.Error("Failed to delete user", )
		return fmt.Errorf("failed to delete user: %w", err)
	}
	logger.Info("User deleted successfully", )
	return nil
}

// GetByEmail mengambil user berdasarkan email
func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("User by email not found", )
			return nil, fmt.Errorf("user not found")
		}
		logger.Error("Failed to get user by email", )
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}
