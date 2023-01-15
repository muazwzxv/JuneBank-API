package app

import (
	"account-service/module/user/domain"
	"account-service/module/user/repository"
	"log"

	"golang.org/x/net/context"
)

type IUserRepository interface {
	Create(ctx context.Context, item *domain.UserData) error
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
