package main

import (
	"account-service/cmd/server"
	"log"
)

func main() {
	server := server.New()

	err := server.Setup()
	if err != nil {
		log.Fatalf("failed to setup server %v", err)
	}

	server.Start()
}
