package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware menambahkan CORS configuration
func CORSMiddleware(allowOrigins string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{allowOrigins}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true

	return cors.New(config)
}
