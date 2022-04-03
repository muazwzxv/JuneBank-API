package entity

import (
	"github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type AccountInterface interface {
	Create(gorm *gorm.DB) error
	GetById(gorm *gorm.DB, id uint) error
}

type Account struct {
	gorm.Model
	Owner    string
	Balance  float64
	Currency string
}

const (
	EUR = "European Dollar"
	USD = "US Dollar"
	SG  = "Singapore Dollar"
	MYR = "Malaysia Ringgit"
)

func (a *Account) ValidateCreate() error {
	return validation.ValidateStruct(a,
		validation.Field(a.Owner, validation.Required),
		validation.Field(a.Balance, validation.Required),
		validation.Field(a.Currency, validation.Required),
	)
}

func (a *Account) Create(gorm *gorm.DB) error {
	if err := gorm.Debug().Create(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Account) GetByID(gorm *gorm.DB, id uint) error {
	if err := gorm.Debug().Where("id = ?", id).First(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Account) DeleteByID(gorm *gorm.DB, id uint) error {
	if err := gorm.Debug().Delete("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
