package server

import (
	"github.com/go-chi/chi/v5"
)

type IServer interface {
	Start()
	SetupServer() error
}

type chiServer struct {
	Mux *chi.Mux
}

var _ IServer = (*chiServer)(nil)

func New() IServer {
	return &chiServer{}
}
