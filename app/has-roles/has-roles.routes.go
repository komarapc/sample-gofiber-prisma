package has_roles

import (
	"github.com/gofiber/fiber/v2"
	"goprisma/db"
)

func Routes(r fiber.Router, p *db.PrismaClient) {
	r.Get("/", func(c *fiber.Ctx) error { return IndexController(c) })
	r.Get("/:id", func(c *fiber.Ctx) error { return ShowController(c) })
	r.Post("/", func(c *fiber.Ctx) error { return StoreController(c) })
	r.Put("/:id", func(c *fiber.Ctx) error { return UpdateController(c) })
	r.Delete("/:id", func(c *fiber.Ctx) error { return DestroyController(c) })
}
