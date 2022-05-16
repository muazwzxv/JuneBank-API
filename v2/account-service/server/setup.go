package server

import (
	"github.com/gofiber/fiber/v2"
	"junebank/v2/account-service/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Setup() {
	logging := log.New(os.Stdout, "JuneBank API V2", log.LstdFlags)

	if err := database.GetGormInstance().Migrate(); err != nil {
		logging.Fatalf("Failed to migrate %v", err)
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   120 * time.Second,
	})

	registerRoutes(app)

	go func() {
		logging.Println("Server starting")
		err := app.Listen(":8080")
		if err != nil {
			logging.Printf("Error starting server: %s \n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig := <-c

	logging.Println("Received terminate request, Grateful shutdown initiated", sig)

	/**
	graceful shutdown goes here
	*/
}

func registerRoutes(app *fiber.App) {
	v2 := app.Group("/api/v2")

	v2.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Alive")
	})
}
