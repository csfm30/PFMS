package expensesource

import (
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type requestDeleteSourceFromName struct {
	UserID uint   `gorm:"not null"`
	Name   string `json:"name"`
}

func DeleteExpenseSourceFromName(c *fiber.Ctx) error {
	db := database.DBConn
	requestDeleteFromName := new(requestDeleteSourceFromName)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	d, _ := strconv.Atoi(userId)
	userIdUint := uint(d)

	//CheckInput
	if err := c.BodyParser(requestDeleteFromName); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	requestDeleteFromName.Name = strings.TrimSpace(requestDeleteFromName.Name)

	if requestDeleteFromName.Name == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}
	var responseDeleteModel modelsPg.ExpenseCategory
	if err := db.Where("name = ? and user_id = ?", requestDeleteFromName.Name, userIdUint).Find(&responseDeleteModel).Delete(modelsPg.ExpenseCategory{}).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	utility.ResetAutoIncrement(db, "expense_categories", "id")

	return utility.ResponseSuccess(c, responseDeleteModel)

}
