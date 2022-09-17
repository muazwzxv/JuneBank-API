package handler

import (
	"junebank_v1/entity"
	"junebank_v1/service"
	"junebank_v1/util"

	"github.com/gofiber/fiber/v2"
)

type transactionHandler struct {
	transactionService service.ITransactionService
}

type ITransactionHandler interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
}

func InitializeTransactionHandler(service service.ITransactionService) ITransactionHandler {
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
	id := util.ParseIdParams(ctx)

	if err := t.transactionService.DeleteByID(id); err != nil {
		return util.BadRequest(ctx, "failed to delete", err)
	}
	return util.Ok(ctx, "transaction deleted", nil)
}
