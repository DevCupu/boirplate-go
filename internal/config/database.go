package config

import (
	"fmt"

	"github.com/DevCupu/boirplate-go/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// InitDatabase menginisialisasi koneksi database
func InitDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		logger.Error("Failed to connect to database")
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Failed to get database instance")
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	logger.Info("Database connected successfully")
	return db, nil
}
