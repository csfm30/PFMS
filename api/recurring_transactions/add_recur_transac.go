package recurringtransactions

import (
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

type requestAddRecurringTransaction struct {
	Type               string     `gorm:"size:10;not null" json:"type"`
	Amount             float64    `gorm:"type:decimal(10,2);not null" json:"amount"`
	StartDate          string     `gorm:"not null" json:"start_date"`
	EndDate            *time.Time `json:"end_date"`
	RecurrencePeriodID *uint      `gorm:"not null" json:"recurren_period_id"`
	SourceID           *uint      `json:"source_id"`
	Description        string     `gorm:"type:text" json:"description"`
}

func AddRecurringTransaction(c *fiber.Ctx) error {
	db := database.DBConn
	requestAddRecurTransac := new(requestAddRecurringTransaction)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	d, _ := strconv.Atoi(userId)
	userIdUint := uint(d)

	//CheckInput
	if err := c.BodyParser(requestAddRecurTransac); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	layout := "2006-01-02" // Layout for just date without time
	parsedDate, _ := time.Parse(layout, requestAddRecurTransac.StartDate)

	requestAddRecurTransac.Description = strings.TrimSpace(requestAddRecurTransac.Description)
	requestAddRecurTransac.Type = strings.TrimSpace(requestAddRecurTransac.Type)

	if requestAddRecurTransac.Type == "" || requestAddRecurTransac.Description == "" || requestAddRecurTransac.Amount == 0 || requestAddRecurTransac.SourceID == nil || requestAddRecurTransac.StartDate == "" || requestAddRecurTransac.RecurrencePeriodID == nil {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	curTransacModel := modelsPg.RecurringTransaction{
		UserId:             userIdUint,
		Type:               requestAddRecurTransac.Type,
		Description:        requestAddRecurTransac.Description,
		Amount:             requestAddRecurTransac.Amount,
		StartDate:          parsedDate,
		SourceID:           requestAddRecurTransac.SourceID,
		RecurrencePeriodID: *requestAddRecurTransac.RecurrencePeriodID,
	}

	if err := db.Create(&curTransacModel).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utility.ResponseSuccess(c, curTransacModel)

}
