package user

import "account-service/app/pkg/core/domain"

func (r *userRepo) GetByID(id uint64) (*domain.User, error) {
	// TODO
	return &domain.User{}, nil
}

func (r *userRepo) Save(domain.CreateUser) error {
	// TODO
	return nil
}
