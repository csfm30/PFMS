package router

import (
	loginregister "pfms/api/login_and_register"

	"github.com/gofiber/fiber/v2"
)

func setRouteLogin(v1 fiber.Router) {
	v1.Post("/register", loginregister.RegisterUser)
	v1.Post("/login_by_username", loginregister.LoginWithUsername)

}
