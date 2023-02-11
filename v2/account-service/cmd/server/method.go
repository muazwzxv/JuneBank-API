package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
	r "github.com/unrolled/render"

	consume "account-service/app/events/consume"
	publish "account-service/app/events/publish"
	userHandler "account-service/app/handlers/user"
	userService "account-service/app/pkg/core/services/user"
	userRepo "account-service/app/pkg/repositories/user"
)

func (s *chiServer) SetupServer() error {
	if s.Mux != nil {
		return errors.New("SERVER ALREADY EXISTS")
	}

	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)
	mux.Use(middleware.AllowContentType("application/json"))
	mux.Use(render.SetContentType(render.ContentTypeJSON))

	s.Mux = mux
	s.Render = r.New(r.Options{
		IndentJSON: true,
	})

	return nil
}

func (s *chiServer) Start() {
	http.ListenAndServe(":3000", s.Mux)
}

func (s *chiServer) SetupHandlers(
	db *sqlx.DB,
	pub *publish.Publisher,
	sub *consume.Subscriber,
) {

	usersrv := userService.New(userRepo.New(db), pub)
	userHandler := userHandler.New(usersrv, s.Render, s.Log)

	s.Mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/user/event", userHandler.TestEvent)
		r.Get("/user/{id}", userHandler.Get)
		r.Post("/user", userHandler.Create)
		r.Put("/user/{id}", userHandler.Update)
		r.Delete("/user/{id}", userHandler.Delete)
	})

}
