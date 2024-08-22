package loginregister

import (
	"errors"
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type requestRegisterUser struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func RegisterUser(c *fiber.Ctx) error {
	db := database.DBConn
	requestRegisterUser := new(requestRegisterUser)

	//CheckInput
	if err := c.BodyParser(requestRegisterUser); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	if requestRegisterUser.Username == "" || requestRegisterUser.Password == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	requestRegisterUser.Username = strings.TrimSpace(requestRegisterUser.Username)
	requestRegisterUser.Password = strings.TrimSpace(requestRegisterUser.Password)

	getCheckUser := modelsPg.User{}
	if err := db.Where("username = ?", requestRegisterUser.Username).Find(&getCheckUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
		}
	}
	if getCheckUser.Username != "" {
		return utility.ResponseError(c, fiber.StatusOK, "username_is_already")
	}

	//Encrypt Password
	passwordEncrypt, err := utility.AESEncrypt(viper.GetString("aes.aes_key"), requestRegisterUser.Password)
	if err != nil {
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	dataUser := modelsPg.User{
		Username:     requestRegisterUser.Username,
		PasswordHash: passwordEncrypt,
		Email:        "not set:" + passwordEncrypt,
		IsSetData:    false,
	}

	//Save

	if err := db.Model(&dataUser).Where("username = ?", dataUser.Username).Debug().FirstOrCreate(&dataUser).Error; err != nil {
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utility.ResponseSuccess(c, dataUser)

}
