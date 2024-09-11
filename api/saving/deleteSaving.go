package saving

import (
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type requestSavingFromeName struct {
	UserID uint   `gorm:"not null"`
	Name   string `json:"name"`
}

func DeleteSavingFromName(c *fiber.Ctx) error {
	db := database.DBConn
	requestDeleteFromName := new(requestSavingFromeName)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	// d, _ := strconv.Atoi(userId)
	// userIdUint := uint(d)

	//CheckInput
	if err := c.BodyParser(requestDeleteFromName); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	requestDeleteFromName.Name = strings.TrimSpace(requestDeleteFromName.Name)

	if requestDeleteFromName.Name == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}
	var responseDeleteModel modelsPg.Saving
	if err := db.Where("name = ? and id = ?", requestDeleteFromName.Name, userId).Find(&responseDeleteModel).Delete(modelsPg.Saving{}).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	if responseDeleteModel.Name == "" {
		return utility.ResponseSuccess(c, "You don't have permission to access or Name don't exist")
	}
	utility.ResetAutoIncrement(db, "savings", "id")

	return utility.ResponseSuccess(c, responseDeleteModel)

}
