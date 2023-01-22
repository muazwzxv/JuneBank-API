package user

import (
	"account-service/app/pkg/core/domain"
)

func (svr *userService) Create(createUser domain.CreateUser) error {
	return svr.userRepository.Save(createUser)
}

func (svr *userService) Get(id uint64) (*domain.User, error) {
	return svr.userRepository.GetByID(id)
}
