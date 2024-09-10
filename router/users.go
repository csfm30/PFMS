package router

import (
	"pfms/api/users"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func setRouteUser(v1 fiber.Router) {
	usersV1 := v1.Group("/users")

	usersV1.Get("/getAllUsers", middleware.AdminAuth(), users.GetAllUsers)
	usersV1.Delete("/deleteUser", middleware.AdminAuth(), users.DeleteUser)
	usersV1.Get("/get-own-user", middleware.AuthJwt(), users.GetOwnUser)

}
