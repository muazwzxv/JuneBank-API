package entity

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Owner    string
	Balance  float64
	Currency string
}

func (a *Account) create(gorm *gorm.DB) error {
	if err := gorm.Debug().Create(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Account) getById(gorm *gorm.DB, id uint) error {
	if err := gorm.Debug().Where("id = ?", id).First(a).Error; err != nil {
		return err
	}
	return nil
}
