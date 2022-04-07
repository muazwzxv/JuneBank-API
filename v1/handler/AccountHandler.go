package handler

import (
	"github.com/gofiber/fiber/v2"
	"junebank/service"
)

type accountHandler struct {
	accountService service.AccountService
}

type AccountHandler interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
	UpdateByID(ctx *fiber.Ctx) error
}

func InitializeAccountHandler(service service.AccountService) AccountHandler {
	return &accountHandler{accountService: service}
}

func (a *accountHandler) Create(ctx *fiber.Ctx) error {
	return nil
}
func (a *accountHandler) GetAll(ctx *fiber.Ctx) error {
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
