package app

import (
	"account-service/module/user/domain"
	"account-service/module/user/repository"
	"log"
)

type IUserRepository interface {
	Create(item *domain.CreateUser) error
}

var _ IUserRepository = (*repository.UserRepository)(nil)

type Module struct {
	log            log.Logger
	UserRepository IUserRepository
}

func NewModule(
	logger log.Logger,
	UserRepository IUserRepository,
) *Module {
	return &Module{
		log:            logger,
		UserRepository: UserRepository,
	}
}
