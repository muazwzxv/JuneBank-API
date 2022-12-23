package api

import "github.com/gofiber/fiber"

type PingAPI interface {
	Ping(ctx *fiber.Ctx)
}
