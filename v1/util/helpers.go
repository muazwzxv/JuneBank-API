package util

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ParseIdParams(ctx *fiber.Ctx) uint {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		_ = BadRequest(ctx, "failed to parse id", err)
	}
	return uint(id)
}
