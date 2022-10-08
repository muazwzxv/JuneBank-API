package handler

import "junebank/v2/account-service/service"

type accountHandler struct {
	accountService service.IAccountService
}

type IAccountHandler interface {
	Create()
}

func CreateAccountHandler(accountService service.IAccountService) IAccountHandler {
	return &accountHandler{accountService}
}

func (h *accountHandler) Create() {
}
