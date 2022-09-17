package entity

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AccountInterface interface {
	ValidateCreate() error
	Create(gorm *gorm.DB) error
	GetById(gorm *gorm.DB, id uint) error
	GetAll(gorm *gorm.DB, ctx *fiber.Ctx) (*[]Account, error)
	DeleteByID(gorm *gorm.DB, id uint) error
}

type Account struct {
	gorm.Model
	Owner    string  `json:"owner"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
	UserId   uint

	User User
}

const (
	EUR = "European Dollar"
	USD = "US Dollar"
	SG  = "Singapore Dollar"
	MYR = "Malaysia Ringgit"
)
