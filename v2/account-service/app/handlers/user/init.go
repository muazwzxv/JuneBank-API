package user

import (
	"account-service/app/pkg/core/ports"
	"log"

	"github.com/unrolled/render"
)

type UserHandler struct {
	userService ports.IUserService
	R           *render.Render
	Log         *log.Logger
}

var _ ports.IUserHandlers = (*UserHandler)(nil)

func New(
	userService ports.IUserService,
	r *render.Render,
	log *log.Logger,
) ports.IUserHandlers {
	return &UserHandler{
		userService: userService,
		R:           r,
		Log:         log,
	}
}
