package router

import (
	"pfms/api/report"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func setRouteReport(v1 fiber.Router) {
	reportV1 := v1.Group("/reports")

	reportV1.Post("/expense-report", middleware.AuthJwt(), report.ExpenseReport)

}
