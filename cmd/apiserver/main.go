package main

import (
	"github.com/ent1k1377/wb_l0/internal/apiserver"
	"log"
)

func main() {
	config := apiserver.NewConfig()
	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
