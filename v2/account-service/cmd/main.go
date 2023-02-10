package main

import (
	"account-service/app/adapter/pg"
	"account-service/app/adapter/rabbitmq"
	consume "account-service/app/events/consume"
	publish "account-service/app/events/publish"
	"account-service/cmd/server"
	"log"
	"os"
)

// TODO: Setup ENV config

// Timestamp in DB is UTC
func main() {
	logger := log.New(os.Stdout, "V2 Account-Service API", log.LstdFlags)
	svr := server.New(logger)

	err := svr.SetupServer()
	if err != nil {
		logger.Fatalf("failed to setup server %v", err)
	}

	db, err := pg.New("postgresql://root:password@localhost:5432/user", logger).
		GetDB()
	if err != nil {
		logger.Fatalf("failed to connect to database: %v", err)
	}

	rbmq := rabbitmq.New(logger)

	pub, err := publish.NewPublisher(rbmq, logger)
	if err != nil {
		logger.Fatalf("failed to connect to rabbitmq %v", err)
	}

	sub, err := consume.NewConsumer(rbmq, logger)
	if err != nil {
		logger.Fatalf("failed to connect to rabbitmq %v", err)
	}

	svr.SetupHandlers(
		db,
		pub,
		sub,
	)

	svr.Start()

}
