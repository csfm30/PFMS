package recurringtransactions

import (
	"errors"
	"pfms/database"
	"pfms/logs"
	"pfms/utility"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type responseGetAllRecurTransactions struct {
	Type        string    `gorm:"size:10;not null" json:"type"` // 'income' or 'expense'
	Amount      float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Source      string    `json:"source"`
	Description string    `gorm:"type:text" json:"description"`
	StartDate   time.Time `gorm:"not null" json:"start_date"`
	PeriodName  string    `gorm:"size:50;not null" json:"period_name"`
}

type reqJoinTable struct {
	Name       string `json:"name"`
	PeriodName string `gorm:"size:50;not null" json:"period_name"`
}

func GetAllRecurTransactions(c *fiber.Ctx) error {
	db := database.DBConn
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	responseGetAllRecurTransactions := []responseGetAllRecurTransactions{}

	// if err := database.CachingCtx().Get("all_recurring_transactions", &responseGetAllRecurTransactions); err == nil {
	// 	logs.Info("Use Redis")
	// 	//* Hit Cache
	// 	return utility.ResponseSuccess(c, responseGetAllRecurTransactions)
	// }

	// allUsers := []modelsPg.Account{}
	if err := db.Table("recurring_transactions").Select("*").Find(&responseGetAllRecurTransactions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ResponseError(c, fiber.StatusOK, "transaction not found")
		}
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	for i, v := range responseGetAllRecurTransactions {
		exportModel := new(reqJoinTable)
		if v.Type == "expense" {
			db.Table("recurring_transactions").
				Select("expense_categories.name, recurrence_periods.period_name").
				Joins("Join expense_categories ON expense_categories.id = recurring_transactions.source_id").
				Joins("Join recurrence_periods ON recurrence_periods.id = recurring_transactions.recurrence_period_id").
				Find(&exportModel)
		} else if v.Type == "income" {
			db.Table("recurring_transactions").
				Select("income_sources.name, recurrence_periods.period_name").
				Joins("Join income_sources ON income_sources.id = recurring_transactions.source_id").
				Joins("Join recurrence_periods ON recurrence_periods.id = recurring_transactions.recurrence_period_id").
				Find(&exportModel)
		}
		responseGetAllRecurTransactions[i].Source = exportModel.Name
		responseGetAllRecurTransactions[i].PeriodName = exportModel.PeriodName
	}

	if err := database.CachingCtx().Set("all_recurring_transactions", &responseGetAllRecurTransactions, time.Minute*2); err != nil {
		logs.Error(err)
	}

	return utility.ResponseSuccess(c, responseGetAllRecurTransactions)

}
