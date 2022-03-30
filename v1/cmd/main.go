package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"junebank/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	database.GetGormInstance().Migrate()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   120 * time.Second,
	})

	logging := log.New(os.Stdout, "JuneBank API", log.LstdFlags)

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

	logging.Println("Received terminate request, Graceful shutdown initiated", sig)

	// Shutdown goes here
}

func setupMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Server().MaxConnsPerIP = 1
}

func setupRoute(app *fiber.App) {
	v1 := app.Group("/api")

	v1.Get("/test", func(ctx *fiber.Ctx) error {
		ctx.Send([]byte("hehe"))
		return nil
	})
}
