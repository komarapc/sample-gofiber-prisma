package has_roles

import (
	"github.com/gofiber/fiber/v2"
	"goprisma/db"
)

func Routes(r fiber.Router, p *db.PrismaClient) {
	r.Get("/user/:userId", func(c *fiber.Ctx) error { return IndexController(c, p) })
	r.Get("/:id", func(c *fiber.Ctx) error { return ShowController(c, p) })
	r.Post("/", func(c *fiber.Ctx) error { return StoreController(c, p) })
	r.Put("/:id", func(c *fiber.Ctx) error { return UpdateController(c, p) })
	r.Delete("/:id", func(c *fiber.Ctx) error { return DestroyController(c, p) })
}
