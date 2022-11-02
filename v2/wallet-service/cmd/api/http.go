package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (a *Application) setup() {
	logging := log.New(os.Stdout, "JuneBank: Wallet Service", log.LstdFlags)

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   120 * time.Second,
	})

	a.fiber = app
	a.log = logging
}

func (a *Application) registerRoutes() {
	v1 := a.fiber.Group("/api/v1")

	v1.Post("/wallet/ping", func(c *fiber.Ctx) error {
		return c.JSON("hehe")
	})
}

func (a *Application) run() {
	go func() {
		a.log.Println("Server starting")
		err := a.fiber.Listen(":8080")
		if err != nil {
			a.log.Printf("Error starting application: %v", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig := <-c

	a.log.Println("Received terminate request, Graceful shutdown initiated", sig)
}
