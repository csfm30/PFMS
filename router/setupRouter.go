package router

import (
	"pfms/api/account"

	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App) {
	apiBackendPrefix := app.Group("/testgo")
	apiRoutes := apiBackendPrefix.Group("/api")
	v1 := apiRoutes.Group("/v1")

	//For test
	v1.Get("/getAllAccount", middleware.AuthJwt(), account.GetAllAccount)

	setRouteLogin(v1)
	setRouteUser(v1)

	setRouteIncome(v1)
	setRouteExpense(v1)

	setRouteTransacion(v1)

	setRouteRecurrencePeriod(v1)

	setRouteRecurringTransaction(v1)

}
