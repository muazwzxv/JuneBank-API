package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"

	userService "account-service/app/pkg/core/services/user"
	userHandler "account-service/app/pkg/handlers/user"
	userRepo "account-service/app/pkg/repositories/user"
)

func (s *chiServer) SetupServer() error {
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

func (s *chiServer) SetupHandlers(db *sqlx.DB) {

	usersrv := userService.New(userRepo.New(db))
	userHandler := userHandler.New(usersrv)

	s.Mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/user/{id}", userHandler.Get)
	})

}
