package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	OwnerName  string    `json:"owner"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Occupation string    `json:"occupation"`
	DOB        time.Time `json:"date_of_birth"`
}

func (a Account) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.OwnerName, validation.Required),
		validation.Field(&a.Phone, validation.Required),
		validation.Field(&a.Email, validation.Required),
		validation.Field(&a.Occupation, validation.Required),
		validation.Field(&a.DOB, validation.Required, validation.Max(time.Now())),
	)
}
