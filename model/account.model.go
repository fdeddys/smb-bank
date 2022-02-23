package model

import "time"

//Account model ...
type Account struct {
	ID           int64     `json:"id" gorm:"column:id"`
	UserName     string    `json:"name" gorm:"column:name"`
	Password     string    `json:"password" gorm:"column:password"`
	Balance      int64     `json:"balance" gorm:"column:balance"`
	LastUpdateBy string    `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"last_update" gorm:"column:last_update"`
}

// TableName ...
func (m *Account) TableName() string {
	return "public.account"
}
