package repository

import (
	"fmt"

	"com.ocbc.smb/constants"
	"com.ocbc.smb/database"
	"com.ocbc.smb/dto"
	"com.ocbc.smb/model"
)

// GetAccount
// Account by id ...
func GetAccount(accId int) (res dto.ContentResponse) {
	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var account model.Account
	err := db.Model(&model.Account{}).Where("id = ?", accId).Find(&account).Error

	if err != nil {
		res.ErrCode = constants.ERR_CODE_02
		res.ErrDesc = constants.ERR_DESC_02_GET_DATA
		res.Contents = err.Error()
		return res
	}
	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_DESC_00_SUCCESS
	res.Contents = account
	return
}

// GetAccountByUsername
// Account by username ...
func GetAccountByUsername(username string) (res dto.ContentResponse) {
	db := database.GetDbCon()
	db.Debug().LogMode(true)

	var account model.Account
	err := db.Model(&model.Account{}).Where("user_name ~* ?", username).Find(&account).Error

	if err != nil {
		res.ErrCode = constants.ERR_CODE_02
		res.ErrDesc = constants.ERR_DESC_02_GET_DATA
		res.Contents = err.Error()
		return res
	}
	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_DESC_00_SUCCESS
	res.Contents = account
	return
}

// RegisterAccount ...
// Register New Account
func RegisterAccount(account model.Account) (res dto.ContentResponse) {
	db := database.GetDbCon()
	db.Debug().LogMode(true)

	if r := db.Save(&account); r.Error != nil {
		res.ErrCode = constants.ERR_CODE_01
		res.ErrDesc = constants.ERR_DESC_01_SAVE_DATABASE
		res.Contents = r.Error
		return res
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_DESC_00_SUCCESS
	res.Contents = fmt.Sprintf("Your user name : " + account.UserName + " , your password : " + account.NewPassword)
	return res
}
