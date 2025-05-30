package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl            string
	Port                   string
	NotificationServiceUrl string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		DatabaseUrl:            os.Getenv("DATABASE_URL"),
		Port:                   os.Getenv("PORT"),
		NotificationServiceUrl: os.Getenv("NOTIFICATION_SERVICE_URL"),
	}
}
