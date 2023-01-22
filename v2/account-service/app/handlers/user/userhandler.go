package user

import (
	"net/http"
)

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.Write([]byte("Get user endpoint here"))
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.Write([]byte("Create user endpoint here"))
}
