package dto

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ============================================================
// STANDARD RESPONSE WRAPPER
// ============================================================

// SuccessResponse wrapper untuk response sukses
type SuccessResponse struct {
	Success   bool        `json:"success"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// ErrorResponse wrapper untuk response error
type ErrorResponse struct {
	Success   bool      `json:"success"`
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Error     string    `json:"error,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// ============================================================
// RESPONSE BUILDERS
// ============================================================

// SuccessJSON mengirim response sukses dengan data
func SuccessJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	response := SuccessResponse{
		Success:   true,
		Code:      statusCode,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
	c.JSON(statusCode, response)
}

// SuccessEmptyJSON mengirim response sukses tanpa data
func SuccessEmptyJSON(c *gin.Context, statusCode int, message string) {
	response := SuccessResponse{
		Success:   true,
		Code:      statusCode,
		Message:   message,
		Timestamp: time.Now(),
	}
	c.JSON(statusCode, response)
}

// ErrorJSON mengirim response error
func ErrorJSON(c *gin.Context, statusCode int, message string, errMsg string) {
	response := ErrorResponse{
		Success:   false,
		Code:      statusCode,
		Message:   message,
		Error:     errMsg,
		Timestamp: time.Now(),
	}
	c.JSON(statusCode, response)
}

// ============================================================
// PAGINATION DTO (Optional)
// ============================================================

// PaginationQuery untuk query pagination
type PaginationQuery struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

// PaginationMeta untuk metadata pagination
type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// PaginatedResponse untuk response dengan pagination
type PaginatedResponse struct {
	Success    bool           `json:"success"`
	Code       int            `json:"code"`
	Message    string         `json:"message"`
	Data       interface{}    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
	Timestamp  time.Time      `json:"timestamp"`
}
