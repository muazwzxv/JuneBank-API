package user

import (
	"account-service/app/api"
	"github.com/gofiber/fiber/v2"
)

type userAPI struct{}

func New() api.IUserAPI {
	return &userAPI{}
}

func (u *userAPI) Create(c *fiber.Ctx) error {
	return c.SendString("We made it here")
}
