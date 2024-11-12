package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUri string
	PORT     string
}

var Envs = Load()

func Load() Config {
	if err := godotenv.Overload(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return Config{
		MongoUri: getEnv("MONGO_URI", "localhost"),
		PORT:     getEnv("PORT", "5001"),
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
