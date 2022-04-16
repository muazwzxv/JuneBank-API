package handler

import (
	"github.com/gofiber/fiber/v2"
	"junebank/service"
)

type transactionHandler struct {
	transactionService service.TransactionService
}

type TransactionHandler interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
	UpdateByID(ctx *fiber.Ctx) error
}

func InitializeTransactionHandler(service service.TransactionService) TransactionHandler {
	return &transactionHandler{transactionService: service}
}

func (t *transactionHandler) Create(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionHandler) GetAll(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionHandler) GetByID(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionHandler) DeleteByID(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionHandler) UpdateByID(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
