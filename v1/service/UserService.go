package service

import (
	"junebank_v1/entity"
	"junebank_v1/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	GetByID(id uint) (*entity.User, error)
}

func InitializeUserService(repository repository.UserRepository) UserService {
	return &userService{userRepository: repository}
}

func (u *userService) GetByID(id uint) (*entity.User, error) {
	if user, err := u.userRepository.GetByID(id); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
