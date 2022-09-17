package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Age           int32  `json:"age"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
	Accounts      []Account
}
