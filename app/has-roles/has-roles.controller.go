package has_roles

import "github.com/gofiber/fiber/v2"

func IndexController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello from has-roles controller!", "action": "index"})
}

func ShowController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello from has-roles controller!", "action": "show"})
}

func StoreController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello from has-roles controller!", "action": "store"})
}

func UpdateController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello from has-roles controller!", "action": "update"})
}

func DestroyController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello from has-roles controller!", "action": "destroy"})
}
