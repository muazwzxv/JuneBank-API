package domain

import "time"

type User struct {
	ID        uint64
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateUser struct {
	FirstName string
	LastName  string
	Email     string
}
