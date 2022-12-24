package main

import (
	"account-service/app/adapter/pg"
	"account-service/cmd/server"
	"log"
)

func main() {
	server := server.New()

	err := server.Setup()
	if err != nil {
		log.Fatalf("failed to setup server %v", err)
	}

	// Start a DB connection
	_ = pg.New("lolololo").
		GetDB()

	server.Start()
}
