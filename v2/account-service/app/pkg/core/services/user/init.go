package user

import "account-service/app/pkg/core/ports"

type userService struct {
	userRepository ports.IUserRepository
}

var _ ports.IUserService = (*userService)(nil)

func New(userRepo ports.IUserRepository) ports.IUserService {
	return &userService{
		userRepository: userRepo,
	}
}
