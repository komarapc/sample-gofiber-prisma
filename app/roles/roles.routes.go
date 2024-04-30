package roles

import (
	"github.com/gofiber/fiber/v2"
	"goprisma/db"
)

func Routes(r fiber.Router, p *db.PrismaClient) {
	r.Get("/", func(c *fiber.Ctx) error { return IndexHandler(c, p) })
	r.Get("/:id", func(c *fiber.Ctx) error { return ShowHandler(c, p) })
	r.Post("/", func(c *fiber.Ctx) error { return StoreHandler(c, p) })
	r.Put("/:id", func(c *fiber.Ctx) error { return UpdateHandler(c, p) })
	r.Delete("/:id", func(c *fiber.Ctx) error { return DeleteHandler(c, p) })
	r.Patch("/:id/recover", func(c *fiber.Ctx) error { return RecoverHandler(c, p) })
}
