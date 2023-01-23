package user

import (
	"account-service/app/pkg/core/domain"
)

type UserResponse struct {
	Data *domain.User `json:"data"`
}

func NewUserReponse(user *domain.User) *UserResponse {
	return &UserResponse{
		Data: user,
	}
}
