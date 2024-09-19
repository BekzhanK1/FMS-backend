package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string
	DBPort             int64
	DBUser             string
	DBPassword         string
	DBName             string
	JWTSecret          string
	JwtExpAccessToken  int64
	JwtExpRefreshToken int64
}

var Envs = Load()

func Load() Config {
	godotenv.Load()

	return Config{
		DBHost:             getEnv("PUBLIC_HOST", "localhost"),
		DBPort:             getEnvAsInt("DB_PORT", 5432),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", "root"),
		DBName:             getEnv("DB_NAME", "fms"),
		JWTSecret:          getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JwtExpAccessToken:  getEnvAsInt("JWT_ACCESS_TOKEN_EXP", 60*5),
		JwtExpRefreshToken: getEnvAsInt("JWT_REFRESH_TOKEN_EXP", 60*60*24*7),
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
