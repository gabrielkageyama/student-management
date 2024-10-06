package main

import (
	"log"

	"github.com/gabrielkageyama/api_teste1/api"
)

func main() {

	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
