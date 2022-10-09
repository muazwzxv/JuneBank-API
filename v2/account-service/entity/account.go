package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Owner    string  `json:"owner"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
	UserId   uint    `json:"user_id"`
}

func (a Account) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Owner, validation.Required),
		validation.Field(&a.Balance, validation.Required),
		validation.Field(&a.Currency, validation.Required),
		validation.Field(&a.UserId, validation.Required),
	)
}

const (
	EUR = "European Dollar"
	USD = "US Dollar"
	SG  = "Singapore Dollar"
	MYR = "Malaysia Ringgit"
)
