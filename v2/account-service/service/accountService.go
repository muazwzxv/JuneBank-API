package service

import (
	"github.com/gofiber/fiber/v2"
	"junebank/v2/account-service/constants"
	"junebank/v2/account-service/entity"
	"junebank/v2/account-service/repository"
)

type accountService struct {
	accountRepository repository.IAccountRepository
}

type IAccountService interface {
	Create(account entity.Account) error
}

func CreateAccountService(repository repository.IAccountRepository) IAccountService {
	return &accountService{accountRepository: repository}
}

func (s *accountService) Create(account entity.Account) error {
	// Validate the payload
	err := account.Validate()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.BAD_REQUEST)
	}

	// Check existing email

	// store in db

	// submit event account created

	return nil
}
