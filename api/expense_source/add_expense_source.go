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

type requestExpenseSource struct {
	UserID      uint   `gorm:"not null"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func AddExpenseSource(c *fiber.Ctx) error {
	db := database.DBConn
	requestAddExpense := new(requestExpenseSource)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	d, _ := strconv.Atoi(userId)
	userIdUint := uint(d)

	//CheckInput
	if err := c.BodyParser(requestAddExpense); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	requestAddExpense.Name = strings.TrimSpace(requestAddExpense.Name)
	requestAddExpense.Description = strings.TrimSpace(requestAddExpense.Description)

	if requestAddExpense.Name == "" || requestAddExpense.Description == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	expenseModel := modelsPg.ExpenseCategory{
		UserId:      userIdUint,
		Name:        requestAddExpense.Name,
		Description: requestAddExpense.Description,
	}
	if err := db.Where("name = ?", requestAddExpense.Name).FirstOrCreate(&expenseModel).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utility.ResponseSuccess(c, expenseModel)

}
