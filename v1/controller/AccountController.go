package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"junebank/database"
)

type AccountRepository struct {
	gorm *gorm.DB
}

func NewAccountController() *AccountRepository {
	db := database.GetGormInstance()
	return &AccountRepository{
		gorm: db.Orm,
	}
}

func (a *AccountRepository) Create(ctx *fiber.Ctx) {

}

func (a *AccountRepository) getAll(ctx *fiber.Ctx) {

}

func (a *AccountRepository) getByID(ctx *fiber.Ctx) {

}

func (a *AccountRepository) deleteByID(ctx *fiber.Ctx) {

}

func (a *AccountRepository) updateByID(ctx *fiber.Ctx) {

}
