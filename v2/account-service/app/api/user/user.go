package user

import (
	"github.com/gofiber/fiber/v2"
)

type userAPI struct{}

func New() IUserAPI {
	return &userAPI{}
}

func (u *userAPI) Create(c *fiber.Ctx) error {
	return c.SendString("We made it here")
}
