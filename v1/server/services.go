package server

import "junebank/service"

type Services struct {
	AccountService     service.AccountService
	TransactionService service.TransactionService
}

func SetupServices(repository *Repositories) *Services {
	account := service.InitializeAccountService(repository.AccountRepository)
	transaction := service.InitializeTransactionService(repository.TransactionRepository)
	return &Services{
		AccountService:     account,
		TransactionService: transaction,
	}
}
