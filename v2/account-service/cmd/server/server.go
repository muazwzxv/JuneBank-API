package server

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

type IServer interface {
	Start() error
	Setup() error
	Routes() error
}

type server struct {
	app *fiber.App
}

func New() IServer {
	var svr server
	return &svr
}

func (s *server) Setup() error {
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

func (s *server) Start() error {
	return s.app.Listen(":3000")
}

func (s *server) Routes() error {
	return nil
}
