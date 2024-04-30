package roles

import (
	"github.com/gofiber/fiber/v2"
	"goprisma/db"
	"goprisma/lib"
)

func requestQueryIndex(c *fiber.Ctx) RoleQueryRequest {
	q := c.Queries()
	perPage := q["per_page"]
	page := q["page"]
	if page == "" {
		page = "1"
	}
	if perPage == "" {
		perPage = "10"
	}
	query := RoleQueryRequest{
		Name:    q["name"],
		Page:    lib.ConvertStringToInt(page),
		PerPage: lib.ConvertStringToInt(perPage),
	}
	return query
}

func IndexHandler(c *fiber.Ctx, prisma *db.PrismaClient) error {
	q := requestQueryIndex(c)
	result := GetAllRoles(q, prisma)
	return c.Status(result.StatusCode).JSON(result)
}

func ShowHandler(c *fiber.Ctx, p *db.PrismaClient) error {
	id := c.Params("id")
	result := GetSingleById(id, p)
	return c.Status(result.StatusCode).JSON(result)
}

func StoreHandler(c *fiber.Ctx, p *db.PrismaClient) error {
	var request RoleRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	result := Store(request, p)
	return c.Status(result.StatusCode).JSON(result)
}

func UpdateHandler(c *fiber.Ctx, p *db.PrismaClient) error {
	id := c.Params("id")
	var request RoleRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	result := Update(id, request, p)
	return c.Status(result.StatusCode).JSON(result)
}

func DeleteHandler(c *fiber.Ctx, p *db.PrismaClient) error {
	id := c.Params("id")
	result := Delete(id, p)
	return c.Status(result.StatusCode).JSON(result)
}

func RecoverHandler(c *fiber.Ctx, p *db.PrismaClient) error {
	id := c.Params("id")
	result := UndeleteRole(id, p)
	return c.Status(result.StatusCode).JSON(result)
}
