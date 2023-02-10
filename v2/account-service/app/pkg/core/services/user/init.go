package user

import (
	pub "account-service/app/events/publish"
	"account-service/app/pkg/core/ports"
)

type userService struct {
	userRepository ports.IUserRepository
	*pub.Publisher
}

var _ ports.IUserService = (*userService)(nil)

func New(userRepo ports.IUserRepository, pub *pub.Publisher) ports.IUserService {
	return &userService{
		userRepository: userRepo,
		Publisher:      pub,
	}
}
