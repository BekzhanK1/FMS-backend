package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBHost     string
	DBPort     int64
	DBUser     string
	DBPassword string
	DBName     string
}

var Envs = Load()

func Load() DBConfig {
	godotenv.Load()

	return DBConfig{
		DBHost:     getEnv("PUBLIC_HOST", "localhost"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBName:     getEnv("DB_NAME", "fms"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
