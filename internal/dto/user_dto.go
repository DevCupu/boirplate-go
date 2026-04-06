package dto

import "time"

// ============================================================
// REQUEST DTOs
// ============================================================

// UserUpdateProfileRequest untuk update profile user (tanpa password)
type UserUpdateProfileRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone" binding:"required,min=10,max=15"`
}

// ChangePasswordRequest untuk change password
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=6"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=100"`
}

// ============================================================
// RESPONSE DTOs
// ============================================================

// UserResponse untuk response user (tidak expose password)
type UserResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	IsActive  bool       `json:"is_active"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// UserListResponse untuk list users
type UserListResponse struct {
	Count int            `json:"count"`
	Users []UserResponse `json:"data"`
}

// ============================================================
// HELPER FUNCTIONS
// ============================================================

// ToUserResponse mengkonversi model User ke DTO UserResponse
func ToUserResponse(id, name, email, phone string, isActive bool, lastLogin *time.Time, createdAt, updatedAt time.Time) UserResponse {
	return UserResponse{
		ID:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		IsActive:  isActive,
		LastLogin: lastLogin,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
