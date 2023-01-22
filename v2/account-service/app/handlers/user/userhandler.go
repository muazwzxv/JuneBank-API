package user

import (
	"net/http"

	"account-service/app/handlers"
	"account-service/app/pkg/core/domain"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")
	userId, err := handlers.StringToUint64(userIdStr)
	if err != nil {
		// returns 404 since URL path could not be converted to uint64
		h.Log.Printf("error parsing id: %d \n %v", userId, err)

		h.R.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid"})
		return
	}

	user, err := h.userService.Get(userId)
	if err != nil {
		h.Log.Printf("error getting user with id: %d \n %v", userId, err)

		h.R.JSON(w, http.StatusNotFound, map[string]string{"error": "user not found"})
		return
	}

	// TODO: Give more metadata in the response
	h.R.JSON(w, http.StatusOK, user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := domain.CreateUser{}
	if err := render.Bind(r, &data); err != nil {
		h.Log.Printf("error binding payload, %v", err)

		h.R.JSON(w, http.StatusBadRequest, map[string]string{"error": "bad request"})
		return
	}

	if err := h.userService.Create(data); err != nil {
		h.Log.Printf("error creating user, %v", err)

		h.R.JSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
		return
	}

	// TODO: Give more metadata in the response
	h.R.JSON(w, http.StatusCreated, map[string]string{"message": "user created"})
}
