package main

import (
	"github.com/ent1k1377/wb_l0/internal/apiserver"
	"log"
)

func main() {
	config := apiserver.NewConfig()
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
