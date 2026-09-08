package config

import (
	"log"

	"taskmanager/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDatabase opens a PostgreSQL connection via GORM and runs
// auto-migrations for all models.
func ConnectDatabase(cfg *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("database connected and migrated successfully")
	return db
}
