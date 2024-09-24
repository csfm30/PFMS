package loginandregister

import (
	"pfms/database"
)

func Logout(userId uint) error {
	db := database.DBConn

	db.Table("users").Where("id = ?", userId).Update("is_login", "false")
	return nil
}
