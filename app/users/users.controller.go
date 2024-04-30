package users

import (
	"goprisma/db"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx, prisma *db.PrismaClient) error {
	result := GetAllUsersService(prisma)
	return c.Status(result.StatusCode).JSON(result)
}

func ShowHandler(c *fiber.Ctx, prisma *db.PrismaClient) error {
	id := c.Params("id")
	result := GetUserByIdService(id, prisma)
	return c.Status(result.StatusCode).JSON(result)
}

func StoreHandler(c *fiber.Ctx, prisma *db.PrismaClient) error {
	var user UserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result := CreateOneService(user, prisma)
	return c.Status(result.StatusCode).JSON(result)
}

func UpdateHandler(c *fiber.Ctx, prisma *db.PrismaClient) error {
	id := c.Params("id")
	var user UserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	result := UpdateOneService(id, user, prisma)
	return c.Status(result.StatusCode).JSON(result)
}

func DeleteHandler(c *fiber.Ctx, prisma *db.PrismaClient) error {
	id := c.Params("id")
	result := DeleteOneService(id, prisma)
	return c.Status(result.StatusCode).JSON(result)
}
