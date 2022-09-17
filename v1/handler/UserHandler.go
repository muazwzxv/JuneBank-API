package handler

import (
	"errors"
	"junebank_v1/service"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService service.UserService
}

type IUserHandler interface {
	GetByID(ctx *fiber.Ctx) error
}

func InitializeUserHandler(service service.UserService) IUserHandler {
	return &userHandler{userService: service}
}

func (u *userHandler) GetByID(ctx *fiber.Ctx) error {
	return errors.New("new error")
}
