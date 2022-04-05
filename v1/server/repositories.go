package server

import (
	"junebank/database"
	"junebank/repository"
)

type Repositories struct {
	AccountRepository repository.AccountRepository
}

func SetupRepositories() *Repositories {
	accountRepository := repository.InitializeAccountRepository(database.GetGormInstance().Orm)
	return &Repositories{
		AccountRepository: accountRepository,
	}
}
