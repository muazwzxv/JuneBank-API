package handler

import (
	"github.com/gofiber/fiber/v2"
	"junebank/v2/account-service/constants"
	"junebank/v2/account-service/entity"
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
	account := new(entity.Account)
	if err := ctx.BodyParser(account); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.PARSE_ERROR)
	}

	return h.accountService.Create(ctx, account)
}
