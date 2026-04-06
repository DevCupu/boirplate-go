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

// UpdateProfile mengupdate profile user
// PUT /api/v1/users/:id
func (u *UserController) UpdateProfile(c *gin.Context) {
	id := c.Param("id")

	var req dto.UserUpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid request body: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user, err := u.service.UpdateProfile(id, &req)
	if err != nil {
		logger.Error("Failed to update profile: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Failed to update profile", err.Error())
		return
	}

	response := dto.ToUserResponse(user.ID, user.Name, user.Email, user.Phone, user.IsActive, user.LastLogin, user.CreatedAt, user.UpdatedAt)
	dto.SuccessJSON(c, http.StatusOK, "Profile updated successfully", response)
}

// ChangePassword mengubah password user
// POST /api/v1/users/:id/change-password
func (u *UserController) ChangePassword(c *gin.Context) {
	id := c.Param("id")

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid request body: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if err := u.service.ChangePassword(id, &req); err != nil {
		logger.Warn("Failed to change password: " + err.Error())
		dto.ErrorJSON(c, http.StatusBadRequest, "Failed to change password", err.Error())
		return
	}

	dto.SuccessEmptyJSON(c, http.StatusOK, "Password changed successfully")
}

// UpdateUser deprecated: use UpdateProfile atau ChangePassword
// Deprecated untuk compatibility, tidak digunakan di routes

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
