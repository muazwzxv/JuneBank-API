package user

import (
	"account-service/app/pkg/core/domain"
)

func (svr *userService) Create(createUser domain.CreateUser) (*domain.User, error) {
	// TODO
	return &domain.User{}, nil
}

func (svr *userService) Get(id uint64) (*domain.User, error) {
	// TODO
	return &domain.User{}, nil
}
