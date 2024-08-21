package utility

import "github.com/gofiber/fiber/v2"

// ResponseError is ResponseError
func ResponseError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}
