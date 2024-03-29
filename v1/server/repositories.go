package server

import (
	"junebank_v1/database"
	"junebank_v1/repository"
)

type Repositories struct {
	AccountRepository     repository.IAccountRepository
	TransactionRepository repository.ITransactionRepository
}

func SetupRepositories() *Repositories {
	account := repository.InitializeAccountRepository(database.GetGormInstance().Orm)
	transaction := repository.InitializeTransactionRepository(database.GetGormInstance().Orm)
	return &Repositories{
		AccountRepository:     account,
		TransactionRepository: transaction,
	}
}
