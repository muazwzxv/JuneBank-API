package server

import "junebank/service"

type Services struct {
	AccountService service.AccountService
}

func SetupServices(repository *Repositories) *Services {
	accountService := service.InitializeAccountService(repository.AccountRepository)
	return &Services{
		AccountService: accountService,
	}
}
