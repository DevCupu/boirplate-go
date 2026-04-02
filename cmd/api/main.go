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
	"github.com/DevCupu/boirplate-go/internal/handler"
	"github.com/DevCupu/boirplate-go/internal/middleware"
	"github.com/DevCupu/boirplate-go/internal/repository"
	"github.com/DevCupu/boirplate-go/internal/service"
	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

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

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)

	// Setup router
	router := setupRouter(cfg, userHandler)

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
func setupRouter(cfg *config.Config, userHandler *handler.UserHandler) *gin.Engine {
	// Set environment
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	// Middleware
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware(cfg.CorsAllowOrigins))
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(gin.Recovery())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"time":   time.Now(),
		})
	})

	// API v1
	v1 := router.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return router
}
