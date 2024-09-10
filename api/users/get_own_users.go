package users

import (
	"errors"
	"pfms/database"
	"pfms/logs"
	"pfms/utility"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type responseGetOwnUser struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func GetOwnUser(c *fiber.Ctx) error {
	db := database.DBConn
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	responseGetOwnUser := responseGetOwnUser{}

	if err := database.CachingCtx().Get("own_user_"+userId, &responseGetOwnUser); err == nil {
		logs.Info("Use Redis")
		//* Hit Cache
		return utility.ResponseSuccess(c, responseGetOwnUser)
	}

	// allUsers := []modelsPg.Account{}
	if err := db.Table("users").Where("id = ?", userId).Select("*").Find(&responseGetOwnUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ResponseError(c, fiber.StatusOK, "username not found")
		}
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := database.CachingCtx().Set("own_user_"+userId, &responseGetOwnUser, time.Minute*2); err != nil {
		logs.Error(err)
	}

	return utility.ResponseSuccess(c, responseGetOwnUser)

}
