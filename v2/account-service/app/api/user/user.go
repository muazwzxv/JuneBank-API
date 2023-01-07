package user

import (
	"account-service/app/usecases"
	"github.com/gofiber/fiber/v2"
)

type userAPI struct {
	createUserUseCase usecases.CreateUserUseCase
}

func New(createUserUseCase usecases.CreateUserUseCase) IUserAPI {
	return &userAPI{createUserUseCase}
}

func (u *userAPI) Create(c *fiber.Ctx) error {
	u.createUserUseCase.CreateUser()

	return c.SendString("We made it here")
}
