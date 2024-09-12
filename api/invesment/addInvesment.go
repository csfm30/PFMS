package invesment

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

type requstInvestment struct {
	Name                string  `gorm:"not null" json:"name"`
	InitialAmount       float64 `gorm:"type:decimal(10,2);not null" json:"initial_amount"`
	CurrentValue        float64 `gorm:"type:decimal(10,2);not null" json:"current_value"`
	MonthlyContribution float64 `gorm:"type:decimal(10,2);not null" json:"monthly_contribution"`
	Return              float64 `gorm:"type:decimal(10,2);not null" json:"return"`
	Notes               string  `gorm:"size:100;not null" json:"notes"`
}

func AddSaving(c *fiber.Ctx) error {
	db := database.DBConn
	reqInvestment := new(requstInvestment)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	d, _ := strconv.Atoi(userId)
	userIdUint := uint(d)
	// d, _ := strconv.Atoi(userId)
	// userIdUint := uint(d)

	//CheckInput
	if err := c.BodyParser(reqInvestment); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	reqInvestment.Name = strings.TrimSpace(reqInvestment.Name)

	if reqInvestment.Name == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	resInvestment := modelsPg.Investment{
		Name:                reqInvestment.Name,
		InitialAmount:       reqInvestment.InitialAmount,
		CurrentValue:        reqInvestment.CurrentValue,
		UserId:              userIdUint,
		MonthlyContribution: reqInvestment.MonthlyContribution,
		Return:              reqInvestment.Return,
		Notes:               reqInvestment.Notes,
	}

	if err := db.Where("name = ?", reqInvestment.Name).FirstOrCreate(&resInvestment).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	return utility.ResponseSuccess(c, resInvestment)

}
