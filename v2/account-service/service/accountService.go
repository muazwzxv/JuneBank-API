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
	Create(ctx *fiber.Ctx, account *entity.Account) error
}

func CreateAccountService(repository repository.IAccountRepository) IAccountService {
	return &accountService{accountRepository: repository}
}

func (s *accountService) Create(ctx *fiber.Ctx, account *entity.Account) error {
	// Validate the payload
	err := account.Validate()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.BAD_REQUEST)
	}

	// check existing email
	if s.accountRepository.IsExistByEmail(account.Email) {
		return fiber.NewError(fiber.StatusBadRequest, constants.EMAIL_IS_USED)
	}

	// check existing phone
	if s.accountRepository.IsExistByPhone(account.Phone) {
		return fiber.NewError(fiber.StatusBadRequest, constants.PHONE_IS_USED)
	}

	// store in db
	if err := s.accountRepository.Create(account); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, constants.SOMETHING_WRONG_HAPPENED)
	}

	//TODO: submit event account created

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": constants.ACCOUNT_CREATED,
	})
}
