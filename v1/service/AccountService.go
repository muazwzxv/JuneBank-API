package service

import (
	"github.com/gofiber/fiber/v2"
	"junebank/entity"
	"junebank/repository"
)

type accountService struct {
	accountRepository repository.AccountRepository
}
type AccountService interface {
	GetAll(ctx *fiber.Ctx) (*[]entity.Account, error)
	GetByID(id uint) (*entity.Account, error)
	Create()
	DeleteByID(id uint) error
}

func InitializeAccountService(repository repository.AccountRepository) AccountService {
	return &accountService{accountRepository: repository}
}

func (a *accountService) GetAll(ctx *fiber.Ctx) (*[]entity.Account, error) {
	if accounts, err := a.accountRepository.GetAll(ctx); err != nil {
		return nil, err
	} else {
		return accounts, nil
	}
}

func (a *accountService) GetByID(id uint) (*entity.Account, error) {
	if account, err := a.accountRepository.GetByID(id); err != nil {
		return nil, err
	} else {
		return account, nil
	}
}

func (a *accountService) Create() {

}

func (a *accountService) DeleteByID(id uint) error {
	if err := a.accountRepository.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
