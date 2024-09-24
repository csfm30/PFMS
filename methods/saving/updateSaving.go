package saving

import (
	"pfms/database"
	"pfms/logs"
	modelsPg "pfms/models/pg"
	"pfms/utility"
)

func UpdateSaving(userid uint, id uint, amount float64) string {
	db := database.DBConn

	resSaving := modelsPg.Saving{}
	err := db.Where("user_id = ? and id = ?", userid, id).Find(&resSaving).Error
	if err != nil {
		logs.Error(err)
		return err.Error()
	}

	resSaving.RemainingAmount = resSaving.RemainingAmount - amount
	resSaving.AmountSaved = resSaving.AmountSaved + amount
	resSaving.CurrentSaving = amount

	utility.ResetAutoIncrement(db, "transactions", "id")

	err = db.Model(&modelsPg.Saving{}).Where("user_id = ? and id = ?", userid, id).Updates(resSaving).Error
	if err != nil {
		logs.Error(err)
		return err.Error()
	}
	return "success"
}
