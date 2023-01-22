package user

import (
	"account-service/app/pkg/core/ports"

	"github.com/unrolled/render"
)

type UserHandler struct {
	userService ports.IUserService
	R           *render.Render
}

var _ ports.IUserHandlers = (*UserHandler)(nil)

func New(
	userService ports.IUserService,
	r *render.Render,
) ports.IUserHandlers {
	return &UserHandler{
		userService: userService,
		R:           r,
	}
}
