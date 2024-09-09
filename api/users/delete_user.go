package users

import (
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func DeleteUser(c *fiber.Ctx) error {
	db := database.DBConn
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	responseDeleteUser := []modelsPg.User{}
	id := c.Query("id")

	if err := db.Where("id = ?", id).Find(&responseDeleteUser).Delete(modelsPg.User{}).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	utility.ResetAutoIncrement(db, "users", "id")

	return utility.ResponseSuccess(c, responseDeleteUser)

}
