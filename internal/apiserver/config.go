package apiserver

import (
	"fmt"
	"os"
)

type Config struct {
	BindAddr    string
	DatabaseURL string
	StanURL     string
}

func NewConfig() *Config {
	databaseURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_CONTAINER_NAME"),
		os.Getenv("DB_CONTAINER_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	stanURL := fmt.Sprintf("nats://%s:%s",
		os.Getenv("STAN_CONTAINER_NAME"),
		os.Getenv("STAN_CONTAINER_PORT"),
	)

	return &Config{
		BindAddr:    ":" + os.Getenv("APP_HOST_PORT"),
		DatabaseURL: databaseURL,
		StanURL:     stanURL,
	}
}
