package router

import (
	"pfms/api/transaction"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func setRouteTransacion(v1 fiber.Router) {
	v1.Post("/add_transaction", middleware.AuthJwt(), transaction.AddTransaction)
	v1.Get("/getAllTransactions", middleware.AdminAuth(), transaction.GetAllTransactions)
	v1.Post("delete_transaction_by_id", middleware.AuthJwt(), transaction.DeleteTransaction)
	v1.Get("/getOwnTransactions", middleware.AuthJwt(), transaction.GetOwnTransactions)
}
