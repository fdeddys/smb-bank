package services

import (
	"com.ocbc.smb/dto"
	"com.ocbc.smb/model"
	"com.ocbc.smb/repository"
)

// AccountService ...
type AccountService struct {
}

// Get Account ...
func (a AccountService) GetAccount(accId int) (res dto.ContentResponse) {

	res.Contents, res.ErrCode, res.ErrDesc = repository.GetAccount(accId)
	return res
}

// Register new Acc
func (a AccountService) RegisterAccount(account *model.Account) dto.ContentResponse {
	var res dto.ContentResponse

	return res
}
