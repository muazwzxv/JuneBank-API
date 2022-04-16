package server

import "junebank/handler"

type Handlers struct {
	AccountHandler     handler.AccountHandler
	TransactionHandler handler.TransactionHandler
}

func SetupHandlers(services *Services) *Handlers {
	account := handler.InitializeAccountHandler(services.AccountService)
	transaction := handler.InitializeTransactionHandler(services.TransactionService)
	return &Handlers{
		AccountHandler:     account,
		TransactionHandler: transaction,
	}
}
