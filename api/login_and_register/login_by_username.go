package loginregister

import (
	"errors"
	"os"
	"pfms/database"
	"pfms/logs"
	"pfms/middleware"
	modelsPg "pfms/models/pg"
	"pfms/utility"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type requestLoginUser struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type responseLoginUser struct {
	Username     string `json:"username" form:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
}

func LoginWithUsername(c *fiber.Ctx) error {
	db := database.DBConn
	requestLoginUser := new(requestLoginUser)

	//CheckInput
	if err := c.BodyParser(requestLoginUser); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	requestLoginUser.Username = strings.TrimSpace(requestLoginUser.Username)
	requestLoginUser.Password = strings.TrimSpace(requestLoginUser.Password)

	if requestLoginUser.Username == "" || requestLoginUser.Password == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	getUser := modelsPg.User{}
	if err := db.Where("username = ?", requestLoginUser.Username).Find(&getUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ResponseError(c, fiber.StatusOK, "username not found")
		}
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	PassworDecrypted, err := utility.AESDecrypt(viper.GetString("aes.aes_key"), getUser.PasswordHash)
	if err != nil {
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	if PassworDecrypted != requestLoginUser.Password {
		return utility.ResponseError(c, fiber.StatusOK, "password is not correct")
	}
	userAccessToken := ""
	userRefreshToken := ""
	if getUser.Role == "user" {
		userAccessToken, userRefreshToken, err = middleware.CreateAuthToken(os.Getenv("ENV"), strconv.Itoa(int(getUser.ID)), getUser.Role)
		if err != nil {
			logs.Error(err)
			return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
		}
	} else if getUser.Role == "admin" {
		userAccessToken, err = middleware.CreateAuthAdminToken(os.Getenv("ENV"), strconv.Itoa(int(getUser.ID)), getUser.Role)
		if err != nil {
			logs.Error(err)
			return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
		}
	} else {
		userAccessToken, userRefreshToken, err = middleware.CreateAuthToken(os.Getenv("ENV"), strconv.Itoa(int(getUser.ID)), getUser.Role)
		if err != nil {
			logs.Error(err)
			return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
		}
		db.Table("users").Where("id = ?", strconv.Itoa(int(getUser.ID))).Update("role", "user")
	}

	//authJwt

	responseLoginUser := responseLoginUser{
		Username:     getUser.Username,
		Email:        getUser.Email,
		Role:         getUser.Role,
		AccessToken:  userAccessToken,
		RefreshToken: userRefreshToken,
	}

	db.Table("users").Where("id = ?", strconv.Itoa(int(getUser.ID))).Update("access_token", userAccessToken).Update("refresh_token", userRefreshToken)

	return utility.ResponseSuccess(c, responseLoginUser)

}
