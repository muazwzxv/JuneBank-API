package application

import (
	"github.com/gofiber/fiber/v2"
	"junebank/v2/account-service/connections"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ApplicationInstance struct {
	Http *fiber.App
	Log  *log.Logger
}

var application *ApplicationInstance

func GetAppInstance() *ApplicationInstance {
	if application == nil {
		panic(1)
	}
	return application
}

func Setup() {
	logging := log.New(os.Stdout, "JuneBank V2: Account Service", log.LstdFlags)
	application.Log = logging
	/**
	Start connections connection
	*/
	if err := connections.GetGormInstance().Migrate(); err != nil {
		application.Log.Fatalf("Failed to migrate %v", err)
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
	application.Http = app
	registerRoutes(application.Http)

	go func() {
		application.Log.Println("Server starting")
		err := application.Http.Listen(":8080")
		if err != nil {
			application.Log.Printf("Error starting application: %s \n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig := <-c

	application.Log.Println("Received terminate request, Grateful shutdown initiated", sig)

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

	routes.Get("/user/:id", handlers.AccountHandler.GetById)
}
