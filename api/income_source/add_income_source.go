package incomesource

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

type requestAddIncome struct {
	UserID      uint   `gorm:"not null"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func AddIncomeSource(c *fiber.Ctx) error {
	db := database.DBConn
	requestAddIncome := new(requestAddIncome)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	d, _ := strconv.Atoi(userId)
	userIdUint := uint(d)

	//CheckInput
	if err := c.BodyParser(requestAddIncome); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	requestAddIncome.Name = strings.TrimSpace(requestAddIncome.Name)
	requestAddIncome.Description = strings.TrimSpace(requestAddIncome.Description)

	if requestAddIncome.Name == "" || requestAddIncome.Description == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	incomeModel := modelsPg.IncomeSource{
		UserId:      userIdUint,
		Name:        requestAddIncome.Name,
		Description: requestAddIncome.Description,
	}

	if err := db.Where("name = ?", requestAddIncome.Name).FirstOrCreate(&incomeModel).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utility.ResponseSuccess(c, incomeModel)

}
