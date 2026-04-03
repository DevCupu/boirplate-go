package controllers

import (
	"net/http"

	"github.com/DevCupu/boirplate-go/internal/dto"
	"github.com/DevCupu/boirplate-go/internal/service"
	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

// CreateUser membuat user baru atau register
// POST /api/v1/users atau /api/v1/auth/register
func (u *UserController) CreateUser(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid request body: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user, err := u.service.CreateUser(req.Name, req.Email, req.Phone, req.Password)
	if err != nil {
		logger.Error("Failed to create user: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Failed to create user", err.Error())
		return
	}

	response := dto.ToUserResponse(user.ID, user.Name, user.Email, user.Phone, user.IsActive, user.LastLogin, user.CreatedAt, user.UpdatedAt)
	dto.SuccessJSON(c, http.StatusCreated, "User created successfully", response)
}

// GetUser mendapatkan user berdasarkan ID
// GET /api/v1/users/:id
func (u *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := u.service.GetUserByID(id)
	if err != nil {
		logger.Warn("User not found: " + id)
		dto.ErrorJSON(c, http.StatusNotFound, "User not found", "")
		return
	}

	response := dto.ToUserResponse(user.ID, user.Name, user.Email, user.Phone, user.IsActive, user.LastLogin, user.CreatedAt, user.UpdatedAt)
	dto.SuccessJSON(c, http.StatusOK, "User retrieved successfully", response)
}

// GetAllUsers mendapatkan semua user
// GET /api/v1/users
func (u *UserController) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		logger.Error("Failed to get users: " + err.Error())
		dto.ErrorJSON(c, http.StatusInternalServerError, "Failed to get users", err.Error())
		return
	}

	userResponses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = dto.ToUserResponse(user.ID, user.Name, user.Email, user.Phone, user.IsActive, user.LastLogin, user.CreatedAt, user.UpdatedAt)
	}

	response := dto.UserListResponse{
		Count: len(userResponses),
		Users: userResponses,
	}
	dto.SuccessJSON(c, http.StatusOK, "Users retrieved successfully", response)
}

// UpdateUser mengupdate user
// PUT /api/v1/users/:id
func (u *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var req dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid request body: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user, err := u.service.UpdateUser(id, req.Name, req.Email, req.Phone, req.Password)
	if err != nil {
		logger.Error("Failed to update user: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Failed to update user", err.Error())
		return
	}

	response := dto.ToUserResponse(user.ID, user.Name, user.Email, user.Phone, user.IsActive, user.LastLogin, user.CreatedAt, user.UpdatedAt)
	dto.SuccessJSON(c, http.StatusOK, "User updated successfully", response)
}

// DeleteUser menghapus user
// DELETE /api/v1/users/:id
func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := u.service.DeleteUser(id); err != nil {
		logger.Error("Failed to delete user: " + err.Error())
		dto.ErrorJSON(c, http.StatusNotFound, "Failed to delete user", err.Error())
		return
	}

	dto.SuccessEmptyJSON(c, http.StatusOK, "User deleted successfully")
}

// Login authenticate user
// POST /api/v1/auth/login
func (u *UserController) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid login request: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	token, user, err := u.service.Login(req.Email, req.Password)
	if err != nil {
		logger.Warn("Login failed: " + err.Error())
		dto.ErrorJSON(c, http.StatusUnauthorized, "Login failed", err.Error())
		return
	}

	userResp := dto.ToUserResponse(user.ID, user.Name, user.Email, user.Phone, user.IsActive, user.LastLogin, user.CreatedAt, user.UpdatedAt)
	response := dto.UserLoginResponse{
		Token:     token,
		ExpiresIn: 24,
		TokenType: "Bearer",
		User:      userResp,
	}

	dto.SuccessJSON(c, http.StatusOK, "Login successful", response)
}
