package main

import (
	"account-service/app/adapter/pg"
	"account-service/cmd/server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	svr := server.New()

	err := svr.Setup()
	if err != nil {
		log.Fatalf("failed to setup server %v", err)
	}

	// Start a DB connection
	// TODO: Log if error is returned
	db, err := pg.New("postgresql://root:password@localhost:5432/user").
		GetDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Setup routes
	if err := svr.Routes(db); err != nil {
		log.Fatalf("failed to setup routes %v", err)
	}

	// Listen from a different goroutine
	go func() {
		if err := svr.Start(); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = svr.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()

	fmt.Println("Fiber was successful shutdown.")
}
