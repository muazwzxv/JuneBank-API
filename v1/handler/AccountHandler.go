package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"junebank/database"
	"junebank/entity"
	"junebank/util"
	"strconv"
)

type AccountHandlerInterface interface {
	Create(ctx *fiber.Ctx) error
	getAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
	UpdateByID(ctx *fiber.Ctx) error
}

type AccountHandler struct {
	gorm *gorm.DB
}

func NewAccountHandler() *AccountHandler {
	db := database.GetGormInstance()
	return &AccountHandler{
		gorm: db.Orm,
	}
}

func (a *AccountHandler) Create(ctx *fiber.Ctx) error {
	account := new(entity.Account)
	if err := ctx.BodyParser(account); err != nil {
		return util.BadRequest(ctx, "Cannot Parse body", err)
	}

	err := account.ValidateCreate()
	if err != nil {
		return util.BadRequest(ctx, "Validation failed", err)
	}

	err = account.Create(a.gorm)
	if err != nil {
		return util.BadRequest(ctx, "Failed to create", err)
	}

	return util.Created(ctx, "Account created", account)
}

func (a *AccountHandler) GetAll(ctx *fiber.Ctx) error {
	// we need pagination here
	account := new(entity.Account)

	accounts, err := account.GetAll(a.gorm, ctx)
	if err != nil {
		return util.BadRequest(ctx, "Can't load accounts", err)
	}

	return util.Ok(ctx, "Accounts found", accounts)
}

func (a *AccountHandler) GetByID(ctx *fiber.Ctx) error {
	account := new(entity.Account)
	id, _ := strconv.Atoi(ctx.Params("id"))

	err := account.GetByID(a.gorm, uint(id))
	if err != nil {
		return util.BadRequest(ctx, "Account not found", err)
	}

	return util.Ok(ctx, "Account found", account)
}

func (a *AccountHandler) DeleteByID(ctx *fiber.Ctx) error {
	return nil
}

func (a *AccountHandler) UpdateByID(ctx *fiber.Ctx) error {
	return nil
}
