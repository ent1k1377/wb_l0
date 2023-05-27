package dr

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_HOST_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
		),
	}
}
