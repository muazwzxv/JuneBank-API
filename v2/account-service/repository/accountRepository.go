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
	IsExistByEmail(email string) bool
	IsExistByPhone(phone string) bool
	GetAccountById(id uint) (*entity.Account, error)
}

func CreateAccountRepository(db *gorm.DB) IAccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) Create(account *entity.Account) error {
	if err := r.db.Debug().Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (r *accountRepository) IsExistByEmail(email string) bool {
	var count int64
	r.db.Debug().Model(&entity.Account{}).Where("email = ?", email).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

func (r *accountRepository) IsExistByPhone(phone string) bool {
	var count int64
	r.db.Debug().Model(&entity.Account{}).Where("phone = ?", phone).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

func (r *accountRepository) GetAccountById(id uint) (*entity.Account, error) {
	account := new(entity.Account)
	if err := r.db.Debug().First(account, id).Error; err != nil {
		return nil, err
	}
	return account, nil
}
