package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DevCupu/boirplate-go/internal/config"
	"github.com/DevCupu/boirplate-go/internal/controllers"
	"github.com/DevCupu/boirplate-go/internal/middleware"
	"github.com/DevCupu/boirplate-go/internal/repository"
	"github.com/DevCupu/boirplate-go/internal/service"
	"github.com/DevCupu/boirplate-go/pkg/auth"
	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize JWT with secret from env
	auth.InitJWT(cfg.JWTSecret)

	// Initialize logger
	if err := logger.InitLogger(cfg.AppEnv); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// Initialize database
	db, err := config.InitDatabase(cfg)
	if err != nil {
		logger.Fatal("Failed to initialize database")
		os.Exit(1)
	}

	// Auto migrate models
	err = db.AutoMigrate(
	// Add your models here
	// &model.User{},
	)
	if err != nil {
		logger.Fatal("Failed to run migrations")
		os.Exit(1)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	authRepo := repository.NewAuthRepository(userRepo) // ← pass userRepo (composition)

	// Initialize services
	authService := service.NewAuthService(authRepo)
	userService := service.NewUserService(userRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)

	// Setup router
	router := setupRouter(cfg, authController, userController)

	// Create server
	srv := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ServerTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.ServerTimeout) * time.Second,
	}

	// Start server in goroutine
	go func() {
		logger.Info("Server starting on port " + cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed")
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown")
		os.Exit(1)
	}

	logger.Info("Server shutdown successfully")
}

// setupRouter mengatur route aplikasi
func setupRouter(cfg *config.Config, authController *controllers.AuthController, userController *controllers.UserController) *gin.Engine {
	// Set environment
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Middleware global
	router.Use(middleware.CORSMiddleware(cfg.CorsAllowOrigins))
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.LoggerMiddleware())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"app":    cfg.AppName,
		})
	})

	// ==================== PUBLIC ROUTES ====================

	// ==================== AUTH ROUTES ====================
	// Auth routes (public - tidak perlu authentication)
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	// ==================== USER ROUTES ====================
	// User routes (public read, protected write)
	users := router.Group("/api/v1/users")
	{
		// Public routes (tidak perlu authentication)
		users.GET("", userController.GetAllUsers)
		users.GET("/:id", userController.GetUser)

		// Protected routes (perlu authentication)
		protected := users.Use(middleware.AuthMiddleware())
		{
			protected.PUT("/:id", userController.UpdateProfile)                   // update profile (name, email, phone)
			protected.POST("/:id/change-password", userController.ChangePassword) // change password
			protected.DELETE("/:id", userController.DeleteUser)
		}
	}

	return router
}
