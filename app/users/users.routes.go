package users

import (
	"github.com/gofiber/fiber/v2"
	"goprisma/db"
)

func Routes(r fiber.Router, prisma *db.PrismaClient) {
	r.Get("/", func(c *fiber.Ctx) error { return IndexHandler(c, prisma) })
	r.Get("/:id", func(ctx *fiber.Ctx) error { return ShowHandler(ctx, prisma) })
	r.Post("/", func(c *fiber.Ctx) error { return StoreHandler(c, prisma) })
	r.Put("/:id", func(c *fiber.Ctx) error { return UpdateHandler(c, prisma) })
	r.Delete("/:id", func(c *fiber.Ctx) error { return DeleteHandler(c, prisma) })
}
