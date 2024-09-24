package loginregister

import (
	"pfms/database"
	"pfms/utility"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Logout(c *fiber.Ctx) error {

	db := database.DBConn
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	d, _ := strconv.Atoi(userId)
	userIdUint := uint(d)
	_ = userIdUint

	db.Table("users").Where("id = ?", userIdUint).Update("is_login", "false")

	return utility.ResponseSuccess(c, "logout success")
}
