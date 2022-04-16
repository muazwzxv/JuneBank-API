package service

import (
	"github.com/gofiber/fiber/v2"
	"junebank/entity"
	"junebank/repository"
)

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

type TransactionService interface {
	GetAll(ctx *fiber.Ctx) (*[]entity.Transaction, error)
	GetByID(id uint) (*entity.Transaction, error)
	Create(transaction *entity.Transaction) error
	DeleteByID(id uint) error
}

func InitializeTransactionService(repository repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository: repository}
}

func (t transactionService) GetAll(ctx *fiber.Ctx) (*[]entity.Transaction, error) {
	if accounts, err := t.transactionRepository.GetAll(ctx); err != nil {
		return nil, err
	} else {
		return accounts, nil
	}
}

func (t transactionService) GetByID(id uint) (*entity.Transaction, error) {
	if account, err := t.transactionRepository.GetById(id); err != nil {
		return nil, err
	} else {
		return account, nil
	}
}

func (t transactionService) Create(transaction *entity.Transaction) error {
	if err := t.transactionRepository.Create(transaction); err != nil {
		return err
	}
	return nil
}

func (t transactionService) DeleteByID(id uint) error {
	if err := t.transactionRepository.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
