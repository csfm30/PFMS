package recurringperiod

import (
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type requestAddecurringPeriod struct {
	PeriodName string `json:"period_name"`
}

func AddRecurringPeriod(c *fiber.Ctx) error {
	db := database.DBConn
	requestAddecurringPeriod := new(requestAddecurringPeriod)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	//CheckInput
	if err := c.BodyParser(requestAddecurringPeriod); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	requestAddecurringPeriod.PeriodName = strings.TrimSpace(requestAddecurringPeriod.PeriodName)

	if requestAddecurringPeriod.PeriodName == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	RecurrencePeriodModel := modelsPg.RecurrencePeriod{
		PeriodName: requestAddecurringPeriod.PeriodName,
	}

	if err := db.Where("period_name = ?", requestAddecurringPeriod.PeriodName).FirstOrCreate(&RecurrencePeriodModel).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utility.ResponseSuccess(c, RecurrencePeriodModel)

}
