package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Environment string
	Server      ServerConfig
	Database    DatabaseConfig
	JWT         JWTConfig
	HospitalAPI HospitalAPIConfig
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port int
}

// DatabaseConfig holds database connection configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	URL      string // Connection string, used for migrations
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret     string
	ExpireTime int // in hours
}

// HospitalAPIConfig holds configuration for external hospital APIs
type HospitalAPIConfig struct {
	HospitalABaseURL string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid server port: %v", err)
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("invalid database port: %v", err)
	}

	jwtExpireTime, err := strconv.Atoi(getEnv("JWT_EXPIRE_TIME", "24"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT expire time: %v", err)
	}

	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "hms")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	// Construct database URL for migrations
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	return &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		Server: ServerConfig{
			Port: port,
		},
		Database: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			DBName:   dbName,
			SSLMode:  dbSSLMode,
			URL:      dbURL,
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your-secret-key"),
			ExpireTime: jwtExpireTime,
		},
		HospitalAPI: HospitalAPIConfig{
			HospitalABaseURL: getEnv("HOSPITAL_A_BASE_URL", "https://hospital-a.api.co.th"),
		},
	}, nil
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
