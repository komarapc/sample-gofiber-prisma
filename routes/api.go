package routes

import (
	"github.com/gofiber/fiber/v2"
	"goprisma/app/users"
	"goprisma/db"
)

func homeHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello World!"})
}

func setupApiRoutes(r *fiber.App, prisma *db.PrismaClient) {
	api := r.Group("/api")
	api.Get("/", homeHandler)
	users.Routes(api.Group("/users"), prisma)
}

func SetupRoutes(r *fiber.App, prisma *db.PrismaClient) {
	setupApiRoutes(r, prisma)
}
