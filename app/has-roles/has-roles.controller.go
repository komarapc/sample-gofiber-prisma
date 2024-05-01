package has_roles

import (
	"goprisma/db"
	"goprisma/lib"

	"github.com/gofiber/fiber/v2"
)

func IndexController(c *fiber.Ctx, p *db.PrismaClient) error {
	userId := c.Params("userId")
	result := GetHasRolesByUserId(userId, p)
	return c.Status(result.StatusCode).JSON(result)
}

func ShowController(c *fiber.Ctx, p *db.PrismaClient) error {
	id := c.Params("id")
	result := GetHasRoleById(id, p)
	return c.Status(result.StatusCode).JSON(result)
}

func StoreController(c *fiber.Ctx, p *db.PrismaClient) error {
	var req HasRolesRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	result := StoreHasRoles(req, p)
	return c.Status(result.StatusCode).JSON(result)
}

func UpdateController(c *fiber.Ctx, p *db.PrismaClient) error {
	id := c.Params("id")
	var request HasRolesRequest
	if err := c.BodyParser(&request); err != nil {
		message := err.Error()
		errResponse := lib.ResponseError(lib.ResponseProps{Code: fiber.StatusUnprocessableEntity, Message: &message})
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errResponse)
	}
	result := UpdateService(id, request, p)
	return c.Status(result.StatusCode).JSON(result)
}

func DestroyController(c *fiber.Ctx, p *db.PrismaClient) error {
	id := c.Params("id")
	result := DestroyService(id, p)
	return c.Status(result.StatusCode).JSON(result)
}
