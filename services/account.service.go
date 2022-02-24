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

	// TODO :
	// Hit Server
	// If success pass Balance to Local Function get account else Throw Error
	currBalance := int64(0)
	return repository.GetAccount(accId, currBalance)
}

// Register new Acc
func (a AccountService) RegisterAccount(account *model.Account) (res dto.ContentResponse) {

	// check if username exist
	if resp := repository.GetAccountByUsername(account.UserName); resp.ErrCode == constants.ERR_CODE_00 {
		res.ErrDesc = constants.ERR_DESC_10_USERNAME_EXIST
		res.ErrCode = constants.ERR_CODE_10
		return
	}
	// TODO :
	// Hit Server
	// If success add to local DB else Throw Error

	// generate PASSWORD
	generatedPlainPassword := utils.GeneratePassword(constants.MAX_LENGTH_PASSWORD)
	hashPassword := sha256.Sum256([]byte(account.UserName + generatedPlainPassword))
	password := hex.EncodeToString(hashPassword[:])

	account.Balance = 0
	account.Password = password
	account.NewPassword = generatedPlainPassword
	account.LastUpdate = time.Now()
	account.LastUpdateBy = account.UserName
	return repository.RegisterAccount(*account)
}

// Login new Acc
func (a AccountService) Login(login *dto.LoginDto) (res dto.ContentResponse) {

	// check if username exist
	resp := repository.GetAccountByUsername(login.Username)
	if resp.ErrCode != constants.ERR_CODE_00 {
		utils.LogInfo("Invalid username : " + login.Username)
		res.ErrDesc = constants.ERR_DESC_53_LOGIN
		res.ErrCode = constants.ERR_CODE_53
		return
	}

	// check if password is valid
	account := resp.Contents.(model.Account)
	if account.Password != login.Password {
		utils.LogInfo("Invalid password :" + login.Username)
		res.ErrDesc = constants.ERR_DESC_53_LOGIN
		res.ErrCode = constants.ERR_CODE_53
		return
	}

	token, err := utils.GenerateToken(account.UserName, account.ID)
	if err != nil {
		res.ErrCode = constants.ERR_CODE_54
		res.ErrDesc = constants.ERR_DESC_54_GENERATE_TOKEN
		res.Contents = err.Error()
		return
	}
	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_DESC_00_SUCCESS
	res.Contents = token

	return
}
