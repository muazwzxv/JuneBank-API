package server

import "junebank/handler"

type Handlers struct {
	AccountHandler handler.AccountHandler
}

func SetupHandlers(services *Services) *Handlers {
	accountHandlers := handler.InitializeAccountHandler(services.AccountService)
	return &Handlers{
		AccountHandler: accountHandlers,
	}
}
