package main

import (
	"account-service/app/adapter/pg"
	"account-service/cmd/server"
	"log"
	"os"
)

// Timestamp in DB is UTC
func main() {
	logger := log.New(os.Stdout, "V2 Account-Service API", log.LstdFlags)
	svr := server.New(logger)

	err := svr.SetupServer()
	if err != nil {
		log.Fatalf("failed to setup server %v", err)
	}

	db, err := pg.New("postgresql://root:password@localhost:5432/user").
		GetDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	svr.SetupHandlers(db)

	svr.Start()
}
