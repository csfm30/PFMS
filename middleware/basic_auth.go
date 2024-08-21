package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"aed":    "redcross",
			"korkla": "@kla7016",
		},
	})
}
