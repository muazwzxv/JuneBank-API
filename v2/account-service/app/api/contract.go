package api

import "github.com/gofiber/fiber/v2"

type IPingAPI interface {
	Ping(ctx *fiber.Ctx)
}

type IUserAPI interface {
	Create(c *fiber.Ctx) error
}
