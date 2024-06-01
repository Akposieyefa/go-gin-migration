package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL   string
	APP_PORT string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		DB_URL:   getEnv("DB_URL", ""),
		APP_PORT: getEnv("APP_PORT", ":9090"),
	}
}

func getEnv(key, fallback string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if _, ok := os.LookupEnv(key); ok {
		return os.Getenv(key)
	}
	return fallback
}
