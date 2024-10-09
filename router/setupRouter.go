package router

import (
	"pfms/api/account"
	"pfms/api/script"

	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App) {
	apiBackendPrefix := app.Group("/testgo")
	apiRoutes := apiBackendPrefix.Group("/api")
	v1 := apiRoutes.Group("/v1")

	//For test
	v1.Get("/getAllAccount", middleware.AuthJwt(), account.GetAllAccount)

	v1.Post("/test-notify", script.TestNotify)
	// v1.Post("/test-notify", script.TestDiscordNotify)

	setRouteLogin(v1)
	setRouteUser(v1)

	setRouteIncome(v1)
	setRouteExpense(v1)

	setRouteTransacion(v1)

	setRouteSaving(v1)

	setRouteReport(v1)

}
