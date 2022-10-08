package service

import "junebank/v2/account-service/repository"

type accountService struct {
	accountRepository repository.IAccountRepository
}

type IAccountService interface {
	Create()
}

func CreateAccountService(repository repository.IAccountRepository) IAccountService {
	return &accountService{accountRepository: repository}
}

func (s *accountService) Create() {
	// Validate the payload

	// Check existing email

	// store in db

	// submit event account created
}
