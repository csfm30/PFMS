package report

import (
	"fmt"
	"pfms/database"
	"pfms/logs"
	"pfms/methods/notify"
	modelsPg "pfms/models/pg"
	"pfms/utility"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type reqExpenseReport struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type responseExpenseReport struct {
	ReqExpenseReport []resExpenseReport
	TotalAmountMonth float64
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

	if err := c.BodyParser(reqExpenseReport); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	var responseExpenseReport responseExpenseReport

	if err := db.Model(&modelsPg.Transaction{}).
		Select("expense_categories.name AS category_name, SUM(transactions.amount) AS total_amount").
		Joins("left join expense_categories ON expense_categories.id = transactions.expense_category_id").
		Where("transactions.type = ? and transactions.user_id = ?", "expense", userId).
		Where("EXTRACT(MONTH FROM transactions.date) = ?", reqExpenseReport.Month).
		Where("EXTRACT(YEAR FROM transactions.date) = ?", reqExpenseReport.Year).
		Group("expense_categories.name").
		Find(&responseExpenseReport.ReqExpenseReport).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	for _, v := range responseExpenseReport.ReqExpenseReport {
		responseExpenseReport.TotalAmountMonth += v.TotalAmount
	}

	if responseExpenseReport.TotalAmountMonth > 20000 {
		monthStr := reqExpenseReport.Month
		monthInt, err := strconv.Atoi(monthStr) // Convert string to integer
		if err != nil {
			fmt.Println("Error converting month:", err)
			return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
		}
		monthName := time.Month(monthInt)
		message := fmt.Sprintf("Your expense of %s is exceed 20,000 baht", &monthName)
		notify.DiscordNotify(message)
	}
	fmt.Println(responseExpenseReport)

	return utility.ResponseSuccess(c, responseExpenseReport)
}
