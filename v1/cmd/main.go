package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"junebank/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "JuneBank API", log.LstdFlags)

	data := database.GetGormInstance()
	fmt.Println(data)

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   120 * time.Second,
	})

	go func() {
		logger.Println("Server starting")
		err := app.Listen(":8080")
		if err != nil {
			logger.Printf("Error starting server: %s \n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig := <-c

	logger.Println("Received terminate request, Graceful shutdown initiated", sig)

	// Shutdown goes here
}

func setupMiddleware() {

}

func setupRoute() {

}
