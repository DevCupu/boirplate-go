package dto

import "time"

// ============================================================
// REQUEST DTOs
// ============================================================

// UserCreateRequest untuk membuat user baru
type UserCreateRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,min=10,max=15"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

// UserUpdateRequest untuk update user
type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,min=10,max=15"`
	Password string `json:"password" binding:"omitempty,min=6,max=100"`
}

// UserLoginRequest untuk login
type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
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

// UserLoginResponse untuk response login
type UserLoginResponse struct {
	Token     string       `json:"token"`
	ExpiresIn int          `json:"expires_in"`
	TokenType string       `json:"token_type"`
	User      UserResponse `json:"user"`
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
