package middleware

import (
	"net/http"

	"github.com/DevCupu/boirplate-go/pkg/logger"

	"github.com/gin-gonic/gin"
)

// APIError adalah struktur error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorHandlerMiddleware menangani error secara global
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Cek apakah ada error
		for _, err := range c.Errors {
			logger.Error("Request error: " + err.Error())

			// Default error response
			code := http.StatusInternalServerError
			message := "Internal Server Error"

			c.JSON(code, gin.H{
				"error": APIError{
					Code:    code,
					Message: message,
				},
			})
			return
		}
	}
}
