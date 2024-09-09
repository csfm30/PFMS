package router

import (
	"pfms/api/account"

	"pfms/api/users"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App) {
	apiBackendPrefix := app.Group("/testgo")
	apiRoutes := apiBackendPrefix.Group("/api")
	v1 := apiRoutes.Group("/v1")

	setRouteLogin(v1)

	v1.Get("/getAllAccount", middleware.AuthJwt(), account.GetAllAccount)
	v1.Get("/getAllUsers", middleware.AuthJwt(), users.GetAllUsers)
	v1.Delete("/deleteUser", middleware.AdminAuth(), users.DeleteUser)

	setRouteIncome(v1)
	setRouteExpense(v1)

	setRouteTransacion(v1)

	setRouteRecurrencePeriod(v1)

	setRouteRecurringTransaction(v1)

}
