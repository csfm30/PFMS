package report

import (
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type reqExpenseReport struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type resExpenseReport struct {
	CategoryName string  `json:"category_name"`
	TotalAmount  float64 `json:"total_amount"`
}

func ExpenseReport(c *fiber.Ctx) error {
	db := database.DBConn
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	reqExpenseReport := new(reqExpenseReport)
	resExpenseReport := []resExpenseReport{}

	if err := c.BodyParser(reqExpenseReport); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	if err := db.Model(&modelsPg.Transaction{}).
		Select("expense_categories.name AS category_name, SUM(transactions.amount) AS total_amount").
		Joins("left join expense_categories ON expense_categories.id = transactions.expense_category_id").
		Where("transactions.type = ? and transactions.user_id = ?", "expense", userId).
		Where("EXTRACT(MONTH FROM transactions.date) = ?", reqExpenseReport.Month).
		Where("EXTRACT(YEAR FROM transactions.date) = ?", reqExpenseReport.Year).
		Group("expense_categories.name").
		Find(&resExpenseReport).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utility.ResponseSuccess(c, resExpenseReport)
}
