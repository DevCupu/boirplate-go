package repository

import (
	"github.com/DevCupu/boirplate-go/internal/model"
)

// AuthRepository interface untuk auth repository
// Thin wrapper yang delegate ke UserRepository (composition)
type AuthRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
	ExistsByEmail(email string) bool
	Update(user *model.User) error
}

// authRepository implementasi dari AuthRepository
type authRepository struct {
	userRepo UserRepository // ← COMPOSITION: delegate ke UserRepository
}

// NewAuthRepository membuat instance baru AuthRepository
func NewAuthRepository(userRepo UserRepository) AuthRepository {
	return &authRepository{userRepo: userRepo}
}

// Create membuat user baru (untuk register) - delegate ke UserRepository
func (r *authRepository) Create(user *model.User) error {
	return r.userRepo.Create(user)
}

// GetByEmail mendapatkan user berdasarkan email (untuk login) - delegate
func (r *authRepository) GetByEmail(email string) (*model.User, error) {
	return r.userRepo.GetByEmail(email)
}

// ExistsByEmail cek apakah email sudah terdaftar - delegate
func (r *authRepository) ExistsByEmail(email string) bool {
	return r.userRepo.ExistsByEmail(email)
}

// Update update user (untuk update last login) - delegate
func (r *authRepository) Update(user *model.User) error {
	return r.userRepo.Update(user)
}
