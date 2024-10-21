package config

import (
	"log"
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
	MailUsername	   string
	MailPassword 	   string
	MailPort		   int64
	MailServer		   string
	CryptoKey		   string
}

var Envs = Load()
var AdminConfig = LoadAdminUserConfig()

func Load() Config {
	if err := godotenv.Overload(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return Config{
		DBHost:             getEnv("PUBLIC_HOST", "localhost"),
		DBPort:             getEnvAsInt("DB_PORT", 5432),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", "root"),
		DBName:             getEnv("DB_NAME", "fms"),
		JWTSecret:          getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JwtExpAccessToken:  getEnvAsInt("JWT_ACCESS_TOKEN_EXP", 60*5),
		JwtExpRefreshToken: getEnvAsInt("JWT_REFRESH_TOKEN_EXP", 60*60*24*7),
		MailUsername:	    getEnv("MAIL_USERNAME", "bkimadieff@gmail.com"),
		MailPassword:	    getEnv("MAIL_PASSWORD", "password"),
		MailPort:		    getEnvAsInt("MAIL_PORT", 587),
		MailServer:		    getEnv("MAIL_SERVER", "smtp.gmail.com"),
		CryptoKey: 		    getEnv("CRYPTO_KEY", "12345678901234567890123456789012"),
	}
}

type AdminUserConfiguration struct {
	Email    string
	Username string
	Phone    string
	Password string
}

func LoadAdminUserConfig() *AdminUserConfiguration {		
	return &AdminUserConfiguration{
		Email:   getEnv("ADMIN_EMAIL", "admin@gmail.com"),
		Username: getEnv("ADMIN_USERNAME", "admin"),
		Phone:   getEnv("ADMIN_PHONE", "+77775556655"),
		Password: getEnv("ADMIN_PASSWORD", "qwerty"),
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
