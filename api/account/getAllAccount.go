package account

import (
	"errors"
	"pfms/database"
	"pfms/utility"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type responseGetAllLogin struct {
	Username     string `json:"username" form:"username"`
	FirstNameTh  string `json:"first_name_th"`
	LastNameTh   string `json:"last_name_th"`
	FirstNameEng string `json:"first_name_eng"`
	LastNameEng  string `json:"last_name_eng"`
	NicknameTh   string `json:"nickname_th"`
	NicknameEng  string `json:"nickname_eng"`
	Email        string `json:"email"`
	MobileNo     string `json:"mobile_no"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
}

func GetAllAccount(c *fiber.Ctx) error {
	db := database.DBConn
	myUser := c.Locals("user").(*jwt.Token)
	claims := myUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	_ = userId

	responseGetAllLogin := []responseGetAllLogin{}

	// allUsers := []modelsPg.Account{}
	if err := db.Table("accounts").Select("username,first_name_th,last_name_th").Find(&responseGetAllLogin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ResponseError(c, fiber.StatusOK, "username not found")
		}
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	// responseLoginUser := responseLoginUser{
	// 	Username:     getUser.Username,
	// 	FirstNameTh:  getUser.FirstNameTh,
	// 	LastNameTh:   getUser.LastNameTh,
	// 	FirstNameEng: getUser.FirstNameEng,
	// 	LastNameEng:  getUser.LastNameEng,
	// 	NicknameTh:   getUser.NicknameTh,
	// 	NicknameEng:  getUser.NicknameEng,
	// 	Email:        getUser.Email,
	// 	MobileNo:     getUser.MobileNo,
	// 	AccessToken:  userAccessToken,
	// 	RefreshToken: userRefreshToken,
	// 	Role:         getUser.Role,
	// }

	return utility.ResponseSuccess(c, responseGetAllLogin)

}
