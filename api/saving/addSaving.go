package saving

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

type requestSaving struct {
	Name          string  `gorm:"not null" json:"name"`
	TargetAmount  float64 `gorm:"type:decimal(10,2);not null" json:"target_amount"`
	CurrentSaving float64 `gorm:"type:decimal(10,2);not null" json:"current_saving"`
	AmountSaved   float64 `gorm:"type:decimal(10,2);not null" json:"amount_saved"`
}

func AddSaving(c *fiber.Ctx) error {
	db := database.DBConn
	reqSaving := new(requestSaving)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId
	d, _ := strconv.Atoi(userId)
	userIdUint := uint(d)

	//CheckInput
	if err := c.BodyParser(reqSaving); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	reqSaving.Name = strings.TrimSpace(reqSaving.Name)

	if reqSaving.Name == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	resSaving := modelsPg.Saving{
		UserId:          userIdUint,
		Name:            reqSaving.Name,
		TargetAmount:    reqSaving.TargetAmount,
		CurrentSaving:   reqSaving.CurrentSaving,
		AmountSaved:     reqSaving.AmountSaved,
		RemainingAmount: reqSaving.TargetAmount - reqSaving.CurrentSaving,
	}

	if err := db.Where("name = ? and user_id = ?", reqSaving.Name, userId).FirstOrCreate(&resSaving).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	return utility.ResponseSuccess(c, resSaving)

}
