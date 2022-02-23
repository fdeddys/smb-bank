package services

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"com.ocbc.smb/constants"
	"com.ocbc.smb/dto"
	"com.ocbc.smb/model"
	"com.ocbc.smb/repository"
	"com.ocbc.smb/utils"
)

// AccountService ...
type AccountService struct {
}

// Get Account ...
func (a AccountService) GetAccount(accId int) (res dto.ContentResponse) {

	return repository.GetAccount(accId)
}

// Register new Acc
func (a AccountService) RegisterAccount(account *model.Account) (res dto.ContentResponse) {

	// check if username exist
	if resp := repository.GetAccountByUsername(account.UserName); resp.ErrCode == constants.ERR_CODE_00 {
		res.ErrDesc = constants.ERR_DESC_10_USERNAME_EXIST
		res.ErrCode = constants.ERR_CODE_10
		return
	}

	// generate PASSWORD
	generatedPlainPassword := utils.GeneratePassword(constants.MAX_LENGTH_PASSWORD)
	hashPassword := sha256.Sum256([]byte(account.UserName + generatedPlainPassword))
	password := hex.EncodeToString(hashPassword[:])

	account.Balance = 0
	account.Password = password
	account.NewPassword = generatedPlainPassword
	account.LastUpdate = time.Now()
	account.LastUpdateBy = dto.CurrUser
	return repository.RegisterAccount(*account)
}
