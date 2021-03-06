package server

import (
	"junebank/database"
	"junebank/repository"
)

type Repositories struct {
	AccountRepository     repository.AccountRepository
	TransactionRepository repository.TransactionRepository
}

func SetupRepositories() *Repositories {
	account := repository.InitializeAccountRepository(database.GetGormInstance().Orm)
	transaction := repository.InitializeTransactionRepository(database.GetGormInstance().Orm)
	return &Repositories{
		AccountRepository:     account,
		TransactionRepository: transaction,
	}
}
