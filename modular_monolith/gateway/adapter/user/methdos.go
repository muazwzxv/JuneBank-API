package user

import "account-service/module/user/domain"

func (adpt *Adapter) CreateUser(user *domain.CreateUser) error {
	return adpt.userClient.CreateUser(user)
}
