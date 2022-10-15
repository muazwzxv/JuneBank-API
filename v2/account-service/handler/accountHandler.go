package handler

import (
	"github.com/gofiber/fiber/v2"
	"junebank/v2/account-service/application/util"
	"junebank/v2/account-service/constants"
	"junebank/v2/account-service/request"
	"junebank/v2/account-service/service"
	"strconv"
)

type accountHandler struct {
	accountService service.IAccountService
}

type IAccountHandler interface {
	Create(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
}

func CreateAccountHandler(accountService service.IAccountService) IAccountHandler {
	return &accountHandler{accountService}
}

func (h *accountHandler) Create(ctx *fiber.Ctx) error {
	req := new(request.CreateAccount)
	if err := ctx.BodyParser(req); err != nil {
		return util.BadRequest(ctx, constants.PARSE_ERROR, err)
	}

	return h.accountService.Create(ctx, req)
}

func (h *accountHandler) GetById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return util.BadRequest(ctx, constants.PARSE_ERROR, err)
	}
	return h.accountService.GetById(ctx, uint(id))
}