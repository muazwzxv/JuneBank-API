package server

import (
	"junebank_v1/caching"
	"junebank_v1/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupServer() {

	logging := log.New(os.Stdout, "JuneBank API", log.LstdFlags)

	if err := database.GetGormInstance().Migrate(); err != nil {
		logging.Fatalf("Failed to migrate %v", err)
	}

	_ = caching.GetRedisInstance()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   120 * time.Second,
	})

	/**
	Instantiate
	- repository
	- service
	- handler
	*/
	handlers := SetupHandlers(SetupServices(SetupRepositories()))
	registerRoutes(app, handlers)

	go func() {
		logging.Println("Server starting")
		err := app.Listen(":8080")
		if err != nil {
			logging.Printf("Error starting application: %s \n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig := <-c

	logging.Println("Received terminate request, Graceful shutdown initiated", sig)

	// Shutdown goes here
}

func registerRoutes(app *fiber.App, handlers *Handlers) {
	v1 := app.Group("/api/v1")

	v1.Post("/account", handlers.AccountHandler.Create)
	v1.Get("/account", handlers.AccountHandler.GetAll)
	v1.Get("/account/:id", handlers.AccountHandler.GetByID)
	v1.Put("/account/:id", handlers.AccountHandler.UpdateByID)
	v1.Delete("/account/:id", handlers.AccountHandler.DeleteByID)

	v1.Post("/transaction", handlers.TransactionHandler.Create)
	v1.Get("/transaction", handlers.TransactionHandler.GetAll)
	v1.Get("/transaction/:id", handlers.TransactionHandler.GetByID)
	v1.Delete("/transaction/:id", handlers.TransactionHandler.DeleteByID)
}
