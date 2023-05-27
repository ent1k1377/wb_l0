package apiserver

import (
	"github.com/ent1k1377/wb_l0/internal/dr"
	"os"
)

type Config struct {
	BindAddr string
	DR       *dr.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":" + os.Getenv("APP_HOST_PORT"),
		DR:       dr.NewConfig(),
	}
}
