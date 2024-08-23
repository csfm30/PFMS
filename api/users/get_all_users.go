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

	if err := database.CachingCtx().Get("all_user", &responseGetAllUsers); err == nil {
		logs.Info("Use Redis")
		//* Hit Cache
		return utility.ResponseSuccess(c, responseGetAllUsers)
	}

	// allUsers := []modelsPg.Account{}
	if err := db.Table("users").Select("*").Find(&responseGetAllUsers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ResponseError(c, fiber.StatusOK, "username not found")
		}
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := database.CachingCtx().Set("all_user", &responseGetAllUsers, time.Minute*2); err != nil {
		logs.Error(err)
	}

	return utility.ResponseSuccess(c, responseGetAllUsers)

}
