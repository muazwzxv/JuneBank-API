package handler

import (
	"github.com/gofiber/fiber/v2"
	"junebank/entity"
	"junebank/service"
	"junebank/util"
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
	transaction := new(entity.Transaction)

	if err := ctx.BodyParser(transaction); err != nil {
		return util.BadRequest(ctx, "failed to parse", err)
	}

	if err := t.transactionService.Create(transaction); err != nil {
		return util.BadRequest(ctx, "failed to create", err)
	} else {
		return util.Created(ctx, "transaction created", transaction)
	}
}

func (t *transactionHandler) GetAll(ctx *fiber.Ctx) error {
	if transactions, err := t.transactionService.GetAll(ctx); err != nil {
		return util.BadRequest(ctx, "failed to fetch transactions", err)
	} else {
		return util.Ok(ctx, "transaction fetched", transactions)
	}
}

func (t *transactionHandler) GetByID(ctx *fiber.Ctx) error {
	id := util.ParseIdParams(ctx)

	if transaction, err := t.transactionService.GetByID(id); err != nil {
		return util.BadRequest(ctx, "failed to fetch transaction", err)
	} else {
		return util.Ok(ctx, "transaction fetched", transaction)
	}
}

func (t *transactionHandler) DeleteByID(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionHandler) UpdateByID(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
