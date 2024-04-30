package routes

import (
	"github.com/gofiber/fiber/v2"
	"goprisma/app/users"
	"goprisma/db"
)

func HomeHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello World!"})
}

func setupApiRoutes(r *fiber.App, prisma *db.PrismaClient) {
	api := r.Group("/api")
	api.Get("/", HomeHandler)
	users.Routes(api.Group("/users"), prisma)
}

func SetupRoutes(r *fiber.App, prisma *db.PrismaClient) {
	setupApiRoutes(r, prisma)
}
