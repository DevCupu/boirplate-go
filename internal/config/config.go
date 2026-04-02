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

		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "boilerplate_go"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		// Server
		ServerTimeout: getEnvInt("SERVER_TIMEOUT", 30),

		// Cors
		CorsAllowOrigins: getEnv("CORS_ALLOW_ORIGINS", "*"),
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
