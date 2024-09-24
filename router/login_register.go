package router

import (
	loginregister "pfms/api/login_and_register"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func setRouteLogin(v1 fiber.Router) {
	v1.Post("/register", loginregister.RegisterUser)
	v1.Post("/login_by_username", loginregister.LoginWithUsername)
	v1.Post("/register_admin", loginregister.RegisterAdmin)

	v1.Post("/logout", middleware.AuthJwt(), loginregister.Logout)
}
