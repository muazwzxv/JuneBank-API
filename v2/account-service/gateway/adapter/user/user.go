package user

import (
	"account-service/module/user/api"
	"account-service/module/user/domain"
)

type UserClient interface {
	CreateUser(user *domain.CreateUser) error
}

type Adapter struct {
	userClient UserClient
}

var _ UserClient = (*api.API)(nil)

func New(userClient UserClient) *Adapter {
	return &Adapter{
		userClient: userClient,
	}
}
