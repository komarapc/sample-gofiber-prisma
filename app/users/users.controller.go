package users

import (
	"github.com/gofiber/fiber/v2"
	"goprisma/db"
	"goprisma/lib"
)

func requestQueryIndex(c *fiber.Ctx) UserQueryRequest {
	q := c.Queries()
	perPage := q["per_page"]
	page := q["page"]
	if page == "" {
		page = "1"
	}
	if perPage == "" {
		perPage = "10"
	}
	query := UserQueryRequest{
		Name:    q["name"],
		Email:   q["email"],
		Page:    lib.ConvertStringToInt(page),
		PerPage: lib.ConvertStringToInt(perPage),
	}
	return query
}

func IndexHandler(c *fiber.Ctx, prisma *db.PrismaClient) error {
	q := requestQueryIndex(c)
	result := GetAllUsersService(q, prisma)
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
