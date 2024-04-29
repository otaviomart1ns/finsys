package config

import (
	"os"
)

type Config struct {
	DBDriver string
	DBSource string
}

func LoadEnv() (*Config, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")

	return &Config{
		DBDriver: dbDriver,
		DBSource: dbSource,
	}, nil
}
