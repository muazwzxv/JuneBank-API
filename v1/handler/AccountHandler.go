package handler

import (
	"github.com/gofiber/fiber/v2"
	"junebank/entity"
	"junebank/service"
	"junebank/util"
	"strconv"
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
	account := new(entity.Account)
	if err := ctx.BodyParser(account); err != nil {
		return util.BadRequest(ctx, "Failed to parse", err)
	}

	if err := a.accountService.Create(account); err != nil {
		return util.BadRequest(ctx, "Failed to create", err)
	} else {
		return util.Created(ctx, "Accounts fetched", account)
	}
}

func (a *accountHandler) GetAll(ctx *fiber.Ctx) error {
	if accounts, err := a.accountService.GetAll(ctx); err != nil {
		return util.BadRequest(ctx, "Cannot fetch accounts", err)
	} else {
		return util.Ok(ctx, "Accounts fetched", accounts)
	}
}

func (a *accountHandler) GetByID(ctx *fiber.Ctx) error {
	if id, err := strconv.Atoi(ctx.Params("id")); err != nil {
		return util.BadRequest(ctx, "Failed to parse id", err)
	} else {

		account, err := a.accountService.GetByID(uint(id))
		if err != nil {
			return util.BadRequest(ctx, "Failed to get account", err)
		}

		return util.Ok(ctx, "Account found", account)
	}
}

func (a *accountHandler) DeleteByID(ctx *fiber.Ctx) error {
	return nil
}

func (a *accountHandler) UpdateByID(ctx *fiber.Ctx) error {
	return nil
}
