package users

import (
	"pfms/database"
	"pfms/logs"
	models "pfms/models/pg"
	"pfms/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type reqUser struct {
	Email string `json:"email"`
}

type resUser struct {
	Username string `gorm:"size:100;not null" json:"username"`
	Email    string `json:"email"`
}

func UpdateProfile(c *fiber.Ctx) error {
	db := database.DBConn

	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)

	reqData := new(reqUser)

	if err := c.BodyParser(reqData); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	if reqData.Email == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	reqData.Email = strings.TrimSpace(reqData.Email)

	resData := resUser{}
	if err := db.Model(&models.User{}).Where("id = ?", userId).Update("email", reqData.Email).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	db.Table("users").Where("id = ?", userId).Select("username", "email").Find(&resData)

	return utility.ResponseSuccess(c, resData)

}
