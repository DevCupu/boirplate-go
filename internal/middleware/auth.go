package middleware

import (
    "net/http"
    "strings"

    "github.com/DevCupu/boirplate-go/pkg/auth"
    "github.com/DevCupu/boirplate-go/pkg/logger"
    "github.com/gin-gonic/gin"
)

// AuthMiddleware middleware untuk protect routes
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            logger.Warn("Missing authorization header")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
            c.Abort()
            return
        }

        // Extract token dari "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            logger.Warn("Invalid authorization header format")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
            c.Abort()
            return
        }

        tokenString := parts[1]

        // Verify token
        claims, err := auth.VerifyToken(tokenString)
        if err != nil {
            logger.Warn("Invalid token: " + err.Error())
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Set user info ke context
        c.Set("user_id", claims.UserID)
        c.Set("email", claims.Email)

        c.Next()
    }
}