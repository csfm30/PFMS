package transaction

import (
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type requestDeleteTransaction struct {
	Id string `json:"id"`
}

func DeleteTransaction(c *fiber.Ctx) error {
	db := database.DBConn
	requestDeleteTransaction := new(requestDeleteTransaction)
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	//CheckInput
	if err := c.BodyParser(requestDeleteTransaction); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	requestDeleteTransaction.Id = strings.TrimSpace(requestDeleteTransaction.Id)

	if requestDeleteTransaction.Id == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}
	var responseDeleteModel modelsPg.Transaction
	if err := db.Where("id = ? and user_id = ?", requestDeleteTransaction.Id, userId).Find(&responseDeleteModel).Delete(modelsPg.Transaction{}).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	utility.ResetAutoIncrement(db, "transactions", "id")

	return utility.ResponseSuccess(c, responseDeleteModel)

}
