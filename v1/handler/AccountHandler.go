package handler

import (
	"github.com/gofiber/fiber/v2"
	"junebank/service"
)

type accountHandler struct {
	service service.AccountService
}

type AccountHandler interface {
	Create(ctx *fiber.Ctx) error
	getAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
	UpdateByID(ctx *fiber.Ctx) error
}

func InitializeAccountHandler(service service.AccountService) AccountHandler {
	return &accountHandler{
		service: service,
	}
}

func (a *accountHandler) Create(ctx *fiber.Ctx) error {
	return nil
}
func (a *accountHandler) getAll(ctx *fiber.Ctx) error {
	return nil
}

func (a *accountHandler) GetByID(ctx *fiber.Ctx) error {
	return nil
}

func (a *accountHandler) DeleteByID(ctx *fiber.Ctx) error {
	return nil
}

func (a *accountHandler) UpdateByID(ctx *fiber.Ctx) error {
	return nil
}
