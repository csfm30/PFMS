package router

import (
	expensesource "pfms/api/expense_source"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func setRouteExpense(v1 fiber.Router) {
	v1.Post("/add_expense", middleware.AuthJwt(), expensesource.AddExpenseSource)
	v1.Post("/delete_expense_by_name", middleware.AuthJwt(), expensesource.DeleteExpenseSourceFromName)

}
