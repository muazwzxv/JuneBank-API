package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *chiServer) Setup() error {
	if s.Mux != nil {
		return errors.New("SERVER ALREADY EXISTS")
	}

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	s.Mux = mux

	return nil
}

func (s *chiServer) Start() {
	http.ListenAndServe(":3000", s.Mux)
}
