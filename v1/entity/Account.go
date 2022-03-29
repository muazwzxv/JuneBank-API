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
