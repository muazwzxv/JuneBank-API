package server

import (
	"account-service/app/api"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"

	createuser "account-service/app/usecases/createuser"
	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app *fiber.App
}

func New() IServer {
	return &fiberServer{}
}

func (s *fiberServer) Setup() error {
	if s.app != nil {
		return errors.New("Server already exists")
	}

	app := fiber.New(fiber.Config{
		StrictRouting: false,
		CaseSensitive: false,
		ReadTimeout:   120 * time.Second,
		WriteTimeout:  60 * time.Second,
		IdleTimeout:   60 * time.Second,
		AppName:       "junebank v2 accounts service",
	})

	s.app = app

	return nil
}

func (s *fiberServer) Shutdown() error {
	return s.app.Shutdown()
}

func (s *fiberServer) Start() error {
	return s.app.Listen(":3000")
}

func (s *fiberServer) Routes(db *sqlx.DB) error {
	// Register usecases here
	handlers := api.New(
		createuser.New(db),
	)

	v1 := s.app.Group("/api/v1")
	v1.Get("/user", handlers.User.Create)

	return nil
}
