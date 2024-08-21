package users

import (
	"errors"
	"pfms/database"
	"pfms/utility"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type responseGetAllUsers struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DBConn
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	responseGetAllUsers := []responseGetAllUsers{}

	// allUsers := []modelsPg.Account{}
	if err := db.Table("users").Select("*").Find(&responseGetAllUsers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ResponseError(c, fiber.StatusOK, "username not found")
		}
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utility.ResponseSuccess(c, responseGetAllUsers)

}
