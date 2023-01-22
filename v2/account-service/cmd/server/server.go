package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/render"
)

type IServer interface {
	Start()
	SetupServer() error
	SetupHandlers(db *sqlx.DB)
}

type chiServer struct {
	Mux    *chi.Mux
	Render *render.Render
}

var _ IServer = (*chiServer)(nil)

func New() IServer {
	return &chiServer{}
}
