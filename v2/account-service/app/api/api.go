package api

import (
	"account-service/app/api/user"
	"account-service/app/usecases"
)

type RestHandlers struct {
	User user.IUserAPI
}

func New(
	createUserUseCase usecases.CreateUserUseCase,
) RestHandlers {
	return RestHandlers{
		User: user.New(createUserUseCase),
	}
}
