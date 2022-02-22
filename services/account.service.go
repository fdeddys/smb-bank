package services

import "com.ocbc.smb/dto"

// AccountService ...
type AccountService struct {
}

// GetBrandFilter ...
func (a AccountService) GetAccount(id int) dto.ContentResponse {

	var res dto.ContentResponse

	// data, errCode, errDesc, err := repository.GetBrandFilter(id)

	// if err != nil {
	// 	res.Contents = nil
	// 	res.ErrCode = "02"
	// 	res.ErrDesc = "Error query data to DB"
	// 	return res
	// }

	// res.Contents = data
	// res.ErrCode = errCode
	// res.ErrDesc = errDesc

	return res
}
