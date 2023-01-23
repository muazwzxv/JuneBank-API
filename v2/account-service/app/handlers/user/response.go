package user

import (
	"account-service/app/pkg/core/domain"
)

const (
	MessageUserCreated = "user created"
	MessageUserUpdated = "user updated"
	MessageUserFound   = "user found"
)

type UserResponse struct {
	Data    *domain.User `json:"data"`
	Message string       `json:"message"`
}

func NewUserReponse(
	user *domain.User,
	msg string,
) *UserResponse {
	return &UserResponse{
		Data:    user,
		Message: msg,
	}
}
