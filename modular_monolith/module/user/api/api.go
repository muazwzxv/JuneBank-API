package api

import (
	"account-service/module/user/app"
	"account-service/module/user/domain"
)

type userApp interface {
	CreateUser(user *domain.CreateUser) error
}

var _ userApp = (*app.Module)(nil)

type API struct {
	user userApp
}

func New(user userApp) *API {
	return &API{
		user: user,
	}
}
