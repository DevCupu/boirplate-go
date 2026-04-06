package dto

// ============================================================
// AUTH REQUEST DTOs
// ============================================================

// RegisterRequest untuk registrasi user baru
type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,min=10,max=15"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

// LoginRequest untuk login user
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// ============================================================
// AUTH RESPONSE DTOs
// ============================================================

// LoginResponse untuk response login dengan token
type LoginResponse struct {
	Token     string       `json:"token"`
	ExpiresIn int          `json:"expires_in"`
	TokenType string       `json:"token_type"`
	User      UserResponse `json:"user"`
}
