package controllers

import (
	"net/http"

	"github.com/DevCupu/boirplate-go/internal/dto"
	"github.com/DevCupu/boirplate-go/internal/service"
	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register membuat user baru
// POST /api/v1/auth/register
func (a *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid request body: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user, err := a.authService.Register(&req)
	if err != nil {
		logger.Error("Failed to register user: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Failed to register user", err.Error())
		return
	}

	response := dto.ToUserResponse(user.ID, user.Name, user.Email, user.Phone, user.IsActive, user.LastLogin, user.CreatedAt, user.UpdatedAt)
	dto.SuccessJSON(c, http.StatusCreated, "User registered successfully", response)
}

// Login authenticate user dan generate token
// POST /api/v1/auth/login
func (a *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid login request: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	token, user, err := a.authService.Login(&req)
	if err != nil {
		logger.Warn("Login failed: " + err.Error())
		dto.ErrorJSON(c, http.StatusUnauthorized, "Login failed", err.Error())
		return
	}

	userResp := dto.ToUserResponse(user.ID, user.Name, user.Email, user.Phone, user.IsActive, user.LastLogin, user.CreatedAt, user.UpdatedAt)
	response := dto.LoginResponse{
		Token:     token,
		ExpiresIn: 24,
		TokenType: "Bearer",
		User:      userResp,
	}

	dto.SuccessJSON(c, http.StatusOK, "Login successful", response)
}
