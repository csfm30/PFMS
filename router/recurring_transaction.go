package router

import (
	recurringtransactions "pfms/api/recurring_transactions"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func setRouteRecurringTransaction(v1 fiber.Router) {
	v1.Post("/add_recurrung_transaction", middleware.AuthJwt(), recurringtransactions.AddRecurringTransaction)
	v1.Get("/get_recurring_transaction", middleware.AuthJwt(), recurringtransactions.GetAllRecurTransactions)

}
