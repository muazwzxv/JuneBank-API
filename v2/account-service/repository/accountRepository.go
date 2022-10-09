package repository

import (
	"gorm.io/gorm"
	"junebank/v2/account-service/entity"
)

type accountRepository struct {
	db *gorm.DB
}

type IAccountRepository interface {
	Create(account *entity.Account) error
}

func CreateAccountRepository(db *gorm.DB) IAccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) Create(account *entity.Account) error {
	return nil
}
