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

func (svr *userService) Update(id uint64, update domain.CreateUser) (*domain.User, error) {
	return svr.userRepository.Update(id, update)
}
