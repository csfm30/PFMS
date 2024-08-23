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

type requestDeleteCurringPeriod struct {
	PeriodName string `json:"period_name"`
}

func DeleteCurringPeriod(c *fiber.Ctx) error {
	db := database.DBConn
	requestDeleteCurringPeriod := new(requestDeleteCurringPeriod)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	//CheckInput
	if err := c.BodyParser(requestDeleteCurringPeriod); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	requestDeleteCurringPeriod.PeriodName = strings.TrimSpace(requestDeleteCurringPeriod.PeriodName)

	if requestDeleteCurringPeriod.PeriodName == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}
	var responseDeleteModel modelsPg.RecurrencePeriod
	if err := db.Where("period_name = ?", requestDeleteCurringPeriod.PeriodName).Find(&responseDeleteModel).Delete(modelsPg.RecurrencePeriod{}).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	utility.ResetAutoIncrement(db, "recurrence_periods", "id")

	return utility.ResponseSuccess(c, responseDeleteModel)

}
