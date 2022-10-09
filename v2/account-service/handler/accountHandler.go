package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"junebank/v2/account-service/service"
)

type accountHandler struct {
	accountService service.IAccountService
}

type IAccountHandler interface {
	Create(ctx *fiber.Ctx) error
}

func CreateAccountHandler(accountService service.IAccountService) IAccountHandler {
	return &accountHandler{accountService}
}

func (h *accountHandler) Create(ctx *fiber.Ctx) error {
	fmt.Println("Reached here")
	return nil
}
