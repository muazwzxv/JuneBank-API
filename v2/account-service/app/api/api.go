package api

import "github.com/gofiber/fiber"

func Ping(ctx *fiber.Ctx) {
	ctx.SendString("I'm Alive")
}
