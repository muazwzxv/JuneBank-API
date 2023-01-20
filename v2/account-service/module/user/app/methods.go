package app

import "account-service/module/user/domain"

func (m *Module) CreateUser(user *domain.CreateUser) error {
	return m.UserRepository.Create(user)
}
