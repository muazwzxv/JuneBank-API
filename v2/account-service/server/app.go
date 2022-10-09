package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"junebank/v2/account-service/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Setup() {
	logging := log.New(os.Stdout, "JuneBank V2: Account Service", log.LstdFlags)

	/**
	Start database connection
	*/
	if err := database.GetGormInstance().Migrate(); err != nil {
		logging.Fatalf("Failed to migrate %v", err)
	}

	/**
	TODO: Start kafka connection
	*/

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
	routes := app.Group("/api/v2/account")
	handlers := SetupHandlers(SetupServices(SetupRepositories()))

	routes.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Alive")
	})

	routes.Post("/user", handlers.AccountHandler.Create)

	routes.Get("/user", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Get all users")
	})

	routes.Get("/user/:id", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fmt.Sprintf("%s %s", "Get user with ID", ctx.Params("id")))
	})
}
