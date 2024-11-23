package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort       string
	DB_HOST          string
	DB_PORT          string
	DB_USER          string
	DB_PASSWORD      string
	DB_NAME          string
	REDIS_HOST       string
	REDIS_PORT       string
	MINIO_ENDPOINT   string
	MINIO_ACCESS_KEY string
	MINIO_SECRET_KEY string
	MINIO_BUCKET_NAME string
	JWT_SECRET       string
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Map environment variables to Config struct
	config := &Config{
		ServerPort:       getEnv("SERVER_PORT", ":8080"),
		DB_HOST:          getEnv("DB_HOST", "postgres_db"),
		DB_PORT:          getEnv("DB_PORT", "5432"),
		DB_USER:          getEnv("DB_USER", "postgres"),
		DB_PASSWORD:      getEnv("DB_PASSWORD", "root"),
		DB_NAME:          getEnv("DB_NAME", "everest"),
		REDIS_HOST:       getEnv("REDIS_HOST", "redis"),
		REDIS_PORT:       getEnv("REDIS_PORT", "6380"),
		MINIO_ENDPOINT:   getEnv("MINIO_ENDPOINT", "minio:9000"),
		MINIO_ACCESS_KEY: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
		MINIO_SECRET_KEY: getEnv("MINIO_SECRET_KEY", "miniopassword"),
		MINIO_BUCKET_NAME: getEnv("MINIO_BUCKET_NAME", "everest"),
		JWT_SECRET:       getEnv("JWT_SECRET", "my_jwt_secret"),
	}

	return config, nil
}

// getEnv is a helper function to get environment variable or return a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
