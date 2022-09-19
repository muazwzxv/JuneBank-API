package middleware

import (
	"encoding/json"
	"junebank_v1/caching"

	"github.com/gofiber/fiber/v2"
)

func veryfyCache(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var data interface{}
	val, err := caching.GetRedisInstance().Cache.Get(caching.Ctx, id).Bytes()
	if err != nil {
		return ctx.Next()
	}
	toJson(val, data)

	return ctx.JSON(fiber.Map{"cached": data})
}

func toJson(val []byte, data interface{}) {
	err := json.Unmarshal(val, data)
	if err != nil {
		panic(err)
	}
}
