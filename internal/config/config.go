package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	RedisURL    string
	JWTSecret   string
}

func getEnv(key, defaultValue string)string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		 Port:        getEnv("PORT", "8080"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost/newsfeed?sslmode=disable"),
        RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
        JWTSecret:   getEnv("JWT_SECRET", "random"),
	}
}
