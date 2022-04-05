package repository

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"junebank/entity"
	"junebank/util"
)

type accountRepository struct {
	gorm *gorm.DB
}

type AccountRepository interface {
	Create(account *entity.Account) error
	GetByID(id uint) (*entity.Account, error)
	GetAll(ctx *fiber.Ctx) (*[]entity.Account, error)
	DeleteByID(id uint) error
}

func InitializeAccountRepository(gorm *gorm.DB) AccountRepository {
	return &accountRepository{gorm}
}

func (r *accountRepository) Create(account *entity.Account) error {
	if err := r.gorm.Debug().Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (r *accountRepository) GetByID(id uint) (*entity.Account, error) {
	account := new(entity.Account)
	if err := r.gorm.Debug().Where("id = ?", id).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *accountRepository) GetAll(ctx *fiber.Ctx) (*[]entity.Account, error) {
	accounts := new([]entity.Account)

	if err := r.gorm.Debug().Scopes(util.Paginate(ctx)).Find(accounts).Error; err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *accountRepository) DeleteByID(id uint) error {
	if err := r.gorm.Debug().Delete("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
