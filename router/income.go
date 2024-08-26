package router

import (
	incomesource "pfms/api/income_source"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func setRouteIncome(v1 fiber.Router) {
	v1.Post("/add_income", middleware.AuthJwt(), incomesource.AddIncomeSource)
	v1.Post("/delete_income_by_name", middleware.AuthJwt(), incomesource.DeleteIncomeSourceFromName)

}
