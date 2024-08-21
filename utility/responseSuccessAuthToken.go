package utility

import "github.com/gofiber/fiber/v2"

// ResponseSuccess is ResponseError
func ResponseSuccessAuthToken(c *fiber.Ctx, data interface{}, token string) error {
	if data != nil {
		return c.JSON(fiber.Map{
			"status":     200,
			"message":    "success",
			"auth_token": "Bearer " + token,
			"data":       data,
		})
	}
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
	})
}
