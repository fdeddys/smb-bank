package model

import "time"

//Account model ...
type Account struct {
	ID           int64     `json:"id" gorm:"column:id"`
	UserName     string    `json:"userName" gorm:"column:user_name"`
	AccountName  string    `json:"accountname" gorm:"column:account_name"`
	Password     string    `json:"-" gorm:"column:password"`
	NewPassword  string    `json:"password" gorm:"-" `
	Balance      int64     `json:"balance" gorm:"column:balance"`
	LastUpdateBy string    `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"last_update" gorm:"column:last_update"`
}

// TableName ...
func (m *Account) TableName() string {
	return "public.account"
}
