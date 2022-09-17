package handler

import (
	"junebank_v1/entity"
	"junebank_v1/service"
	"junebank_v1/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type accountHandler struct {
	accountService service.IAccountService
}

type IAccountHandler interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
	UpdateByID(ctx *fiber.Ctx) error
}

func InitializeAccountHandler(service service.IAccountService) IAccountHandler {
	return &accountHandler{accountService: service}
}

func (a *accountHandler) Create(ctx *fiber.Ctx) error {
	account := new(entity.Account)
	if err := ctx.BodyParser(account); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed to parse")
	}

	if err := a.accountService.Create(account); err != nil {
		return fiber.NewError(fiber.ErrConflict.Code, "failed to parse")
	} else {
		return util.Created(ctx, "accounts created", account)
	}
}

func (a *accountHandler) GetAll(ctx *fiber.Ctx) error {
	if accounts, err := a.accountService.GetAll(ctx); err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, "cannot fetch accounts")
	} else {
		return util.Ok(ctx, "accounts fetched", accounts)
	}
}

func (a *accountHandler) GetByID(ctx *fiber.Ctx) error {
	id := util.ParseIdParams(ctx)

	if account, err := a.accountService.GetByID(id); err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, "cannot fetch account")
	} else {
		return util.Ok(ctx, "account found", account)
	}
}

func (a *accountHandler) DeleteByID(ctx *fiber.Ctx) error {
	id := util.ParseIdParams(ctx)

	if _, err := a.accountService.GetByID(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.ErrNotFound.Code, "account not found")
		}
		return fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	if err := a.accountService.DeleteByID(id); err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, "failed to delete account")
	}

	return util.AcceptedNoContent(ctx, "account deleted")
}

func (a *accountHandler) UpdateByID(ctx *fiber.Ctx) error {
	return nil
}
