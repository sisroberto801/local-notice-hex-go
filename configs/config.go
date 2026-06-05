package configs

import (
	"context"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	JWTSecret   string
	Environment string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/notice_db?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func (c *Config) ConnectDB() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), c.DatabaseURL)
}
