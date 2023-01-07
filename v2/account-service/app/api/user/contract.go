package user

import "github.com/gofiber/fiber/v2"

type IUserAPI interface {
	Create(c *fiber.Ctx) error
}
