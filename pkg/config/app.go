package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName string
	DBHost string
	DBPass string
	DBUser string
}

var (
	config Config
)

func GetConfigurations() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config = Config{
		DBName: os.Getenv("DB_NAME"),
		DBHost: os.Getenv("DB_HOST"),
		DBPass: os.Getenv("DB_PASS"),
		DBUser: os.Getenv("DB_USER"),
	}
	return &config
}
