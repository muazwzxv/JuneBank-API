package util

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Return `controller.Error()` with `StatusForbidden`(403)
func Forbidden(ctx *fiber.Ctx, msg string, err error) error {
	return Error(ctx, msg, err, http.StatusForbidden)
}

// Return `controller.Error()` with `StatusConflict`(409)
func Conflict(ctx *fiber.Ctx, msg string, err error) error {
	return Error(ctx, msg, err, http.StatusConflict)
}

// Return `controller.Error()` with `StatusNotFound`(404)
func NotFound(ctx *fiber.Ctx, msg string, err error) error {
	return Error(ctx, msg, err, http.StatusNotFound)
}

// Return `controller.Error()` with `StatusBadRequest`(400)
func BadRequest(ctx *fiber.Ctx, msg string, err error) error {
	return Error(ctx, msg, err, http.StatusBadRequest)
}

// Return `controller.Success()` with `StatusCreated`(201) status code
func Created(ctx *fiber.Ctx, msg string, data interface{}) error {
	return Success(ctx, msg, data, http.StatusCreated)
}

// Return `controller.Success()` with `StatusOk`(20o) status code
func Ok(ctx *fiber.Ctx, msg string, data interface{}) error {
	return Success(ctx, msg, data, http.StatusOK)
}

// Return response with http status code `status`, and JSON message with
// success = true, message = `msg`, and data = `data`
func Success(ctx *fiber.Ctx, msg string, data interface{}, status int) error {
	return response(ctx, status, fiber.Map{
		"success": true,
		"message": msg,
		"data":    data,
	})
}

// Return response with http status code `status`, and JSON message with
// success = false, message = `msg`, and data = `e`
func Error(ctx *fiber.Ctx, msg string, e error, status int) error {
	return response(ctx, status, fiber.Map{
		"success": false,
		"message": msg,
		"data":    e.Error(),
	})
}

// Send a response with status and message.
func response(ctx *fiber.Ctx, status int, msg fiber.Map) error {
	return ctx.Status(status).JSON(msg)
}
