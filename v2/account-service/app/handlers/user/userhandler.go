package user

import (
	"net/http"

	"account-service/app/handlers"

	"github.com/go-chi/chi/v5"
)

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")
	userId, err := handlers.StringToUint64(userIdStr)
	if err != nil {
		// returns 404 since URL path could not be converted to uint64
		h.R.JSON(w, http.StatusNotFound, map[string]string{"error": "invalid"})
		return
	}

	user, err := h.userService.Get(userId)
	if err != nil {
		h.R.JSON(w, http.StatusNotFound, map[string]string{"error": "user not found"})
		return
	}

	// TODO: Give more metadata in the response
	h.R.JSON(w, http.StatusOK, user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.Write([]byte("Create user endpoint here"))
}
