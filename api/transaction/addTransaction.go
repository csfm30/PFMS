package transaction

import (
	"fmt"
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type requestAddTransaction struct {
	Type              string  `gorm:"size:10;not null" json:"type"`
	Amount            float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
	IncomeSourceID    *uint   `json:"income_source_id"`    // Foreign key to IncomeSource, nullable
	ExpenseCategoryID *uint   `json:"expense_category_id"` // Foreign key to ExpenseCategory, nullable
	Description       string  `gorm:"type:text" json:"description"`
}

func AddTransaction(c *fiber.Ctx) error {
	db := database.DBConn
	requestAddTransaction := new(requestAddTransaction)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	d, _ := strconv.Atoi(userId)
	userIdUint := uint(d)

	//CheckInput
	if err := c.BodyParser(requestAddTransaction); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}
	if requestAddTransaction.IncomeSourceID == nil || requestAddTransaction.ExpenseCategoryID == nil {
		fmt.Println("SourceID is nil")
	} else {
		fmt.Printf("SourceID: %d\n", *requestAddTransaction.IncomeSourceID)
	}

	requestAddTransaction.Description = strings.TrimSpace(requestAddTransaction.Description)
	requestAddTransaction.Type = strings.TrimSpace(requestAddTransaction.Type)

	if requestAddTransaction.Type == "" || requestAddTransaction.Description == "" || requestAddTransaction.Amount == 0 || requestAddTransaction.IncomeSourceID == nil || requestAddTransaction.ExpenseCategoryID == nil {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	transactionModel := modelsPg.Transaction{
		UserId:            userIdUint,
		Type:              requestAddTransaction.Type,
		Description:       requestAddTransaction.Description,
		IncomeSourceID:    requestAddTransaction.IncomeSourceID,
		ExpenseCategoryID: requestAddTransaction.ExpenseCategoryID,
		Amount:            requestAddTransaction.Amount,
		Date:              time.Now(),
	}
	if err := db.Create(&transactionModel).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utility.ResponseSuccess(c, transactionModel)

}
