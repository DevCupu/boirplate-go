package repository

import (
	"github.com/DevCupu/boirplate-go/internal/model"
	"gorm.io/gorm"
)

// UserRepository interface untuk user repository
type UserRepository interface {
	Create(user *model.User) error
	GetByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]model.User, error)
	Update(user *model.User) error
	Delete(id string) error
	ExistsByEmail(email string) bool
}

// userRepository implementasi dari UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository membuat instance baru UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create membuat user baru
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// GetByID mendapatkan user berdasarkan ID
func (r *userRepository) GetByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail mendapatkan user berdasarkan email
func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll mendapatkan semua user
func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Where("is_active = ?", true).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Update mengupdate user
func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete menghapus user
func (r *userRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.User{}).Error
}

// ExistsByEmail cek apakah email sudah terdaftar
func (r *userRepository) ExistsByEmail(email string) bool {
	var count int64
	r.db.Model(&model.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
