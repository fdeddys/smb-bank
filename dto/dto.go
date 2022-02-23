package dto

// Current User ...
var CurrUser string

type ContentResponse struct {
	ErrCode  string      `json:"errCode"`
	ErrDesc  string      `json:"errDesc"`
	Contents interface{} `json:"contents"`
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
