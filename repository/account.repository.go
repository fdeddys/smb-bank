package repository

import (
	"com.ocbc.smb/constants"
	"com.ocbc.smb/database"
	"com.ocbc.smb/model"
)

// GetBrandLike ...
func GetAccount(accId int) (interface{}, string, string) {
	db := database.GetDbCon()

	var account model.Account
	err := db.Model(&model.Account{}).Where("id = ?", accId).Find(&account).Error

	if err != nil {
		return err.Error(), constants.ERR_CODE_02, constants.ERR_DESC_02
	}
	return account, constants.ERR_CODE_00, constants.ERR_DESC_00
}
