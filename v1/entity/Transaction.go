package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	SenderId   string
	ReceiverId string
	Amount     float64

	SenderAccount   Account `gorm:"foreignKey:ReceiverId"`
	ReceiverAccount Account `gorm:"foreignKey:SenderId"`
}
