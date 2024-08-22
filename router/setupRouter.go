package router

import (
	"pfms/api/account"
	incomesource "pfms/api/income_source"

	expensesource "pfms/api/expense_source"
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

	v1.Post("/add_income", middleware.AuthJwt(), incomesource.AddIncomeSource)
	v1.Post("/delete_income_by_name", middleware.AuthJwt(), incomesource.DeleteIncomeSourceFromName)

	v1.Post("/add_expense", middleware.AuthJwt(), expensesource.AddExpenseSource)
	v1.Post("/delete_expense_by_name", middleware.AuthJwt(), expensesource.DeleteExpenseSourceFromName)

}
