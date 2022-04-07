package handler

import (
	"github.com/gofiber/fiber/v2"
	"junebank/entity"
	"junebank/service"
	"junebank/util"
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
	id := util.ParseIdParams(ctx)

	if account, err := a.accountService.GetByID(id); err != nil {
		return util.BadRequest(ctx, "Failed to get account", err)
	} else {
		return util.Ok(ctx, "Account found", account)
	}
}

func (a *accountHandler) DeleteByID(ctx *fiber.Ctx) error {
	id := util.ParseIdParams(ctx)
	if err := a.accountService.DeleteByID(id); err != nil {
		return util.BadRequest(ctx, "failed to delete account", err)
	}
	return util.AcceptedNoContent(ctx, "account deleted")
}

func (a *accountHandler) UpdateByID(ctx *fiber.Ctx) error {
	return nil
}
