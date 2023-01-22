package user

import "account-service/app/pkg/core/ports"

type UserHandler struct {
	userService ports.IUserService
}

var _ ports.IUserHandlers = (*UserHandler)(nil)

func New(userService ports.IUserService) ports.IUserHandlers {
	return &UserHandler{
		userService: userService,
	}
}
