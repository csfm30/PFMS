package router

import (
	recurringperiod "pfms/api/recurring_period"
	"pfms/middleware"

	"github.com/gofiber/fiber/v2"
)

func setRouteRecurrencePeriod(v1 fiber.Router) {
	v1.Post("/add_recurring_period", middleware.AuthJwt(), recurringperiod.AddRecurringPeriod)
	v1.Post("/delete_recurring_period", middleware.AuthJwt(), recurringperiod.DeleteCurringPeriod)

}
