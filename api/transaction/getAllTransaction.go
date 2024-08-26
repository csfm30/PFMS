package transaction

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

type responseGetAllTransactions struct {
	Type        string    `gorm:"size:10;not null" json:"type"` // 'income' or 'expense'
	Amount      float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Source      string    `json:"source"`
	Description string    `gorm:"type:text" json:"description"`
	Date        time.Time `gorm:"not null"`
}

func GetAllTransactions(c *fiber.Ctx) error {
	db := database.DBConn
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	responseGetAllTransactions := []responseGetAllTransactions{}

	if err := database.CachingCtx().Get("all_transactions", &responseGetAllTransactions); err == nil {
		logs.Info("Use Redis")
		//* Hit Cache
		return utility.ResponseSuccess(c, responseGetAllTransactions)
	}

	// allUsers := []modelsPg.Account{}
	if err := db.Table("transactions").Select("*").Find(&responseGetAllTransactions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ResponseError(c, fiber.StatusOK, "transaction not found")
		}
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	for i, v := range responseGetAllTransactions {
		exportModel := ""
		if v.Type == "expense" {
			db.Table("transactions").
				Select("expense_categories.name").
				Joins("Join expense_categories ON expense_categories.id = transactions.source_id").
				Find(&exportModel)
		} else if v.Type == "income" {
			db.Table("transactions").
				Select("income_sources.name").
				Joins("Join income_sources ON income_sources.id = transactions.source_id").
				Find(&exportModel)
		}
		responseGetAllTransactions[i].Source = exportModel
	}

	if err := database.CachingCtx().Set("all_transactions", &responseGetAllTransactions, time.Minute*1); err != nil {
		logs.Error(err)
	}

	return utility.ResponseSuccess(c, responseGetAllTransactions)

}
