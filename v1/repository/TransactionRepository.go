package repository

import (
	"junebank_v1/entity"
	"junebank_v1/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type transactionRepository struct {
	gorm *gorm.DB
}

type ITransactionRepository interface {
	Create(transaction *entity.Transaction) error
	GetById(id uint) (*entity.Transaction, error)
	GetAll(ctx *fiber.Ctx) (*[]entity.Transaction, error)
	DeleteByID(id uint) error
}

func InitializeTransactionRepository(gorm *gorm.DB) ITransactionRepository {
	return &transactionRepository{gorm}
}

func (t transactionRepository) Create(transaction *entity.Transaction) error {
	if err := t.gorm.Debug().Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

func (t transactionRepository) GetById(id uint) (*entity.Transaction, error) {
	transaction := new(entity.Transaction)
	if err := t.gorm.Debug().Where("id = ?", id).First(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t transactionRepository) GetAll(ctx *fiber.Ctx) (*[]entity.Transaction, error) {
	transactions := new([]entity.Transaction)
	if err := t.gorm.Debug().Scopes(util.Paginate(ctx)).Find(transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t transactionRepository) DeleteByID(id uint) error {
	if err := t.gorm.Debug().Delete("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
