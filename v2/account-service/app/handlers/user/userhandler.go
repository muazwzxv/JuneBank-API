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
		h.R.JSON(w, http.StatusBadRequest, handlers.ErrInvalidRequest(err))
		return
	}

	userData, err := h.userService.Get(userId)
	if err != nil {
		h.Log.Printf("error getting user with id: %d \n %v", userId, err)
		h.R.JSON(w, http.StatusNotFound, handlers.ErrNotFound(err))
		return
	}

	h.R.JSON(w, http.StatusOK, NewUserReponse(userData, MessageUserFound))
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := domain.CreateUser{}
	if err := render.Bind(r, &data); err != nil {
		h.Log.Printf("error binding payload, %v", err)
		h.R.JSON(w, http.StatusBadRequest, handlers.ErrInvalidRequest(err))
		return
	}

	if err := h.userService.Create(data); err != nil {
		h.Log.Printf("error creating user, %v", err)
		h.R.JSON(w, http.StatusInternalServerError, handlers.ErrInternalServer(err))
		return
	}

	h.R.JSON(w, http.StatusCreated, NewUserReponse(nil, MessageUserCreated))
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {

	userIdStr := chi.URLParam(r, "id")
	userId, err := handlers.StringToUint64(userIdStr)
	if err != nil {
		h.Log.Printf("error parsing id: %d \n %v", userId, err)
		h.R.JSON(w, http.StatusBadRequest, handlers.ErrInvalidRequest(err))
		return
	}

	toUpdate := domain.CreateUser{}
	if err := render.Bind(r, &toUpdate); err != nil {
		h.Log.Printf("error binding payload, %v", err)
		h.R.JSON(w, http.StatusBadRequest, handlers.ErrInvalidRequest(err))
		return
	}

	updated, err := h.userService.Update(userId, toUpdate)
	if err != nil {
		h.R.JSON(w, http.StatusInternalServerError, handlers.ErrInternalServer(err))
		return
	}

	h.R.JSON(w, http.StatusOK, NewUserReponse(updated, MessageUserUpdated))
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

	userIdStr := chi.URLParam(r, "id")
	userId, err := handlers.StringToUint64(userIdStr)
	if err != nil {
		h.Log.Printf("error parsing id: %d \n %v", userId, err)
		h.R.JSON(w, http.StatusBadRequest, handlers.ErrInvalidRequest(err))
		return
	}

	if isExist, err := h.userService.IsUserExist(userId); err != nil {
		h.Log.Printf("h.user.delete %v", err)
		h.R.JSON(w, http.StatusNotFound, handlers.ErrNotFound(err))
		return
	} else {
		if !isExist {
			h.Log.Printf("h.user.delete %v", err)
			h.R.JSON(w, http.StatusNotFound, ErrUserDoesNotExist())
			return
		}
	}

	deletedUser, err := h.userService.Delete(userId)
	if err != nil {
		h.Log.Printf("h.user.delete %v", err)
		h.R.JSON(w, http.StatusInternalServerError, handlers.ErrInternalServer(err))
		return
	}

	h.R.JSON(w, http.StatusOK, NewUserReponse(deletedUser, MessageUserDeleted))
}

func (h *UserHandler) TestEvent(w http.ResponseWriter, r *http.Request) {
	if err := h.userService.TriggerUserEvent(); err != nil {
		h.R.JSON(w, http.StatusInternalServerError, handlers.ErrInternalServer(err))
	}

	h.R.JSON(w, http.StatusOK, &UserResponse{Message: "event triggered"})

}
