package repository

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"junebank/entity"
)

type transactionRepository struct {
	gorm *gorm.DB
}

type TransactionRepository interface {
	Create() error
	GetById(id uint) (*entity.Transaction, error)
	GetAll(ctx *fiber.Ctx) (*[]entity.Transaction, error)
	DeleteByID(id uint) error
}

func InitializeTransactionRepository(gorm *gorm.DB) TransactionRepository {
	return &transactionRepository{gorm}
}

func (t transactionRepository) Create() error {
	//TODO implement me
	panic("implement me")
}

func (t transactionRepository) GetById(id uint) (*entity.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t transactionRepository) GetAll(ctx *fiber.Ctx) (*[]entity.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t transactionRepository) DeleteByID(id uint) error {
	//TODO implement me
	panic("implement me")
}
