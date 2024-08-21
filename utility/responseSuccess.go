package utility

import "github.com/gofiber/fiber/v2"

// ResponseSuccess is ResponseError
func ResponseSuccess(c *fiber.Ctx, data interface{}) error {
	if data != nil {
		return c.JSON(fiber.Map{
			"status":  200,
			"data":    data,
			"message": "success",
		})
	}
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
	})
}
