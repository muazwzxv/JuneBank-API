package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"junebank/database"
	"junebank/entity"
	"junebank/util"
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

	return nil
}

func (a *AccountHandler) GetAll(ctx *fiber.Ctx) error {
	// we need pagination here
	return nil
}

func (a *AccountHandler) GetByID(ctx *fiber.Ctx) error {
	return nil
}

func (a *AccountHandler) DeleteByID(ctx *fiber.Ctx) error {
	return nil
}

func (a *AccountHandler) UpdateByID(ctx *fiber.Ctx) error {
	return nil
}
