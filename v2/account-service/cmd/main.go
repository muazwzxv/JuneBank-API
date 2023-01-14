package main

import (
	"account-service/app/adapter/pg"
	"account-service/cmd/server"
	"log"
)

func main() {
	svr := server.New()

	err := svr.Setup()
	if err != nil {
		log.Fatalf("failed to setup server %v", err)
	}

	// Start a DB connection
	// TODO: Log if error is returned
	_, err = pg.New("postgresql://root:password@localhost:5432/user").
		GetDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	svr.Start()
}
