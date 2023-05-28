package dr

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	conn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_CONTAINER_NAME"),
		os.Getenv("DB_CONTAINER_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	return &Config{
		DatabaseURL: conn,
	}
}
