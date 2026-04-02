package handler

import (
	"net/http"
	"strconv"

	"github.com/DevCupu/boirplate-go/internal/model"
	"github.com/DevCupu/boirplate-go/internal/service"
	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

// UserHandler handle HTTP request untuk user
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler membuat instance UserHandler baru
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser handle POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		logger.Error("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    user,
	})
}

// GetUser handle GET /users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		logger.Warn("User not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User retrieved successfully",
		"data":    user,
	})
}

// GetAllUsers handle GET /users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	limit := 10
	offset := 0

	// Parse query parameter
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	if o := c.Query("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	users, total, err := h.userService.GetAllUsers(limit, offset)
	if err != nil {
		logger.Error("Failed to get users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved successfully",
		"data":    users,
		"pagination": gin.H{
			"total":  total,
			"limit":  limit,
			"offset": offset,
		},
	})
}

// UpdateUser handle PUT /users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var req model.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.userService.UpdateUser(id, &req)
	if err != nil {
		logger.Warn("Failed to update user")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"data":    user,
	})
}

// DeleteUser handle DELETE /users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.userService.DeleteUser(id)
	if err != nil {
		logger.Warn("Failed to delete user")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
