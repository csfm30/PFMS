package router

import (
	"pfms/api/account"
	incomesource "pfms/api/income_source"
	recurringperiod "pfms/api/recurring_period"
	recurringtransactions "pfms/api/recurring_transactions"
	"pfms/api/transaction"

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

	v1.Post("/add_transaction", middleware.AuthJwt(), transaction.AddTransaction)
	v1.Get("/getAllTransactions", middleware.AuthJwt(), transaction.GetAllTransactions)

	v1.Post("/add_recurring_period", middleware.AuthJwt(), recurringperiod.AddRecurringPeriod)
	v1.Post("/delete_recurring_period", middleware.AuthJwt(), recurringperiod.DeleteCurringPeriod)

	v1.Post("/add_recurrung_transaction", middleware.AuthJwt(), recurringtransactions.AddRecurringTransaction)

}
