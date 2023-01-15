package repository

import (
	"account-service/module/user/domain"
	"context"
)

func (r *UserRepository) Create(ctx context.Context, item *domain.UserData) error {
	return nil
}
