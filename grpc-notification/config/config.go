package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl        string
	Port               string
	EmailProviderUrl   string
	EmailProviderToken string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		DatabaseUrl:        os.Getenv("DATABASE_URL"),
		Port:               os.Getenv("PORT"),
		EmailProviderUrl:   os.Getenv("EMAIL_PROVIDER_URL"),
		EmailProviderToken: os.Getenv("EMAIL_PROVIDER_TOKEN"),
	}
}
