package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(max int, expiresInSecond int) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: time.Duration(expiresInSecond) * time.Second,
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"message": "Too many requests. Please try again later."})
		}})
}
