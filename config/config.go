package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver     string
	DBSource     string
	ServerAdress string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	return &Config{
		DBDriver:     dbDriver,
		DBSource:     dbSource,
		ServerAdress: serverAddress,
	}, nil
}
