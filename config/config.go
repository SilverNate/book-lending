package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Port      string
	DBDsn     string
	JWTSecret string
	JWTIssuer string
}

func LoadEnv() *EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("env file not found, using system environment")
	}

	return &EnvConfig{
		Port:      getEnv("PORT", "8080"),
		DBDsn:     getEnv("DB_DSN", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
		JWTIssuer: getEnv("JWT_ISSUER", ""),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
