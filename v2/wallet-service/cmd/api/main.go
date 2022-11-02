package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	_ "github.com/jackc/pgx/v5"
)

type Application struct {
	fiber *fiber.App
	log   *log.Logger
}

func main() {

	var app Application

	app.setup()
	app.registerRoutes()

}
