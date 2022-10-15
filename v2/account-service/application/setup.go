package application

import (
	"junebank/v2/account-service/connections"
	"junebank/v2/account-service/handler"
	"junebank/v2/account-service/repository"
	"junebank/v2/account-service/service"
)

/*
Instantiate Repositories
*/
type Repositories struct {
	AccountRepository repository.IAccountRepository
}

func SetupRepositories() *Repositories {
	account := repository.CreateAccountRepository(connections.GetGormInstance().Orm)
	return &Repositories{
		AccountRepository: account,
	}
}

/*
Instantiate Services
*/
type Services struct {
	AccountService service.IAccountService
}

func SetupServices(repositories *Repositories) *Services {
	account := service.CreateAccountService(repositories.AccountRepository)
	return &Services{
		AccountService: account,
	}
}

/*
Instantiate Handlers
*/
type Handlers struct {
	AccountHandler handler.IAccountHandler
}

func SetupHandlers(services *Services) *Handlers {
	account := handler.CreateAccountHandler(services.AccountService)
	return &Handlers{
		AccountHandler: account,
	}
}
