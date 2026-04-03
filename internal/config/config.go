package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DevCupu/boirplate-go/pkg/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	// App
	AppName string
	AppEnv  string
	AppPort string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Server
	ServerTimeout int

	// Cors
	CorsAllowOrigins string

	// JWT
	JWTSecret string
}

// LoadConfig membaca konfigurasi dari environment
func LoadConfig() *Config {
	// Load .env file (opsional)
	_ = godotenv.Load()

	cfg := &Config{
		// App
		AppName: getEnv("APP_NAME", "boilerplate-go"),
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),

		// Database - No hardcoded passwords, use from .env
		DBHost:     getEnvRequired("DB_HOST"),
		DBPort:     getEnvRequired("DB_PORT"),
		DBUser:     getEnvRequired("DB_USER"),
		DBPassword: getEnvRequired("DB_PASSWORD"),
		DBName:     getEnvRequired("DB_NAME"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		// Server
		ServerTimeout: getEnvInt("SERVER_TIMEOUT", 30),

		// Cors
		CorsAllowOrigins: getEnv("CORS_ALLOW_ORIGINS", "*"),

		// JWT - Load from env (required in production)
		JWTSecret: getEnvRequired("JWT_SECRET"),
	}

	logger.Info("Configuration loaded successfully")
	return cfg
}

// GetDSN mengembalikan database connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBSSLMode,
	)
}

// getEnv membaca environment variable dengan default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// getEnvInt membaca environment variable sebagai integer
func getEnvInt(key string, defaultVal int) int {
	val := getEnv(key, "")
	if val == "" {
		return defaultVal
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		logger.Warn("Failed to convert env to int")
		return defaultVal
	}
	return intVal
}

// getEnvRequired membaca environment variable yang WAJIB ada
func getEnvRequired(key string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}

	logger.Fatal("Required environment variable not found: " + key)
	os.Exit(1)
	return ""
}
