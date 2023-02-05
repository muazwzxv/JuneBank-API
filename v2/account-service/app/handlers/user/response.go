package user

import (
	"account-service/app/handlers"
	"account-service/app/pkg/core/domain"
)

const (
	MessageUserCreated = "user created"
	MessageUserUpdated = "user updated"
	MessageUserFound   = "user found"
	MessageUserDeleted = "user deleted"
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

// Error Response

func ErrUserDoesNotExist() *handlers.ErrResponse {
	return &handlers.ErrResponse{
		StatusText: "request failed",
		ErrorText:  "user does not exists",
	}
}
