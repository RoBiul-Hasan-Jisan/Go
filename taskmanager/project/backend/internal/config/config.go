package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all environment-based configuration for the application.
type Config struct {
	Port            string
	DatabaseURL     string
	JWTSecret       string
	JWTExpiryHours  int
	FrontendURL     string
}

// Load reads the .env file (if present) and returns a populated Config.
// In production, real environment variables are used instead of .env.
func Load() *Config {
	// It's fine if .env doesn't exist (e.g. in production where env vars
	// are injected by the platform), so we only log a soft warning.
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on system environment variables")
	}

	expiry, err := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "72"))
	if err != nil {
		expiry = 72
	}

	cfg := &Config{
		Port:           getEnv("PORT", "8080"),
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpiryHours: expiry,
		FrontendURL:    getEnv("FRONTEND_URL", "http://localhost:3000"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required but was not set")
	}
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required but was not set")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
