package main

import (
	"account-service/app/adapter/pg"
	"account-service/cmd/server"
	"log"
)

func main() {
	svr := server.New()

	err := svr.SetupServer()
	if err != nil {
		log.Fatalf("failed to setup server %v", err)
	}

	_, err = pg.New("postgresql://root:password@localhost:5432/user").
		GetDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	svr.Start()
}
