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
	GetAll(ctx fiber.Ctx) (*[]entity.Transaction, error)
	GetByID(id uint) (*entity.Transaction, error)
	Create(transaction *entity.Transaction) error
	DeleteByID(id uint) error
}

func InitializeTransactionService(repository repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository: repository}
}

func (t transactionService) GetAll(ctx fiber.Ctx) (*[]entity.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t transactionService) GetByID(id uint) (*entity.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t transactionService) Create(transaction *entity.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t transactionService) DeleteByID(id uint) error {
	//TODO implement me
	panic("implement me")
}
