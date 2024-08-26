package middleware

import (
	"pfms/logs"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func CreateAuthAdminToken(env string, userId string, role string) (accessToken string, err error) {
	// Create AccessToken
	token := jwt.New(jwt.SigningMethodHS256)
	refId := uuid.NewV4()

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["type"] = "access"
	claims["env"] = env
	claims["user_id"] = userId
	claims["ref_id"] = refId
	claims["role"] = role

	//sessionTimeAccess := 1440 // Access อายุ 1 วัน
	sessionTimeAccess := 7200 // Access อายุ 5 วัน
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(sessionTimeAccess)).Unix()

	// Generate encoded token and send it as response.
	accessToken, err = token.SignedString([]byte(viper.GetString("auth.admin")))
	if err != nil {
		logs.Error(err)
		return "create_token_fail", err
	}

	return accessToken, err
}

func AdminAuth() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningKey:   []byte(viper.GetString("auth.admin")),
		ErrorHandler: jwtError,
	})
}

//func AdminAuth() fiber.Handler {
//	return basicauth.New(basicauth.Config{
//		Users: map[string]string{
//			viper.GetString("master.username"): viper.GetString("master.password"),
//		},
//	})
//}
