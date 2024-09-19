package saving

import (
	"fmt"
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type requestUpdateSaving struct {
	Name          string  `gorm:"not null" json:"name"`
	CurrentSaving float64 `gorm:"type:decimal(10,2);not null" json:"current_saving"`
}

func UpdateSaving(c *fiber.Ctx) error {
	db := database.DBConn
	reqSaving := new(requestUpdateSaving)
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

	pointerCurrentSaving := &reqSaving.CurrentSaving

	if reqSaving.Name == "" && pointerCurrentSaving == nil {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}
	fmt.Println("d")

	resSaving := modelsPg.Saving{}
	err := db.Where("user_id = ? and name = ?", userIdUint, reqSaving.Name).Find(&resSaving).Error
	if err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	resSaving.RemainingAmount = resSaving.RemainingAmount - reqSaving.CurrentSaving
	resSaving.AmountSaved = resSaving.CurrentSaving + reqSaving.CurrentSaving
	resSaving.CurrentSaving = reqSaving.CurrentSaving

	err = db.Model(&modelsPg.Saving{}).Where("user_id = ? and name = ?", userIdUint, reqSaving.Name).Updates(resSaving).Error
	if err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utility.ResponseSuccess(c, resSaving)

}
