package entity

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Owner    string  `json:"owner"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
	UserId   uint    `json:"user_id"`
}

const (
	EUR = "European Dollar"
	USD = "US Dollar"
	SG  = "Singapore Dollar"
	MYR = "Malaysia Ringgit"
)
