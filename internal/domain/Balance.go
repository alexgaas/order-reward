package domain

import "database/sql"

type Balance struct {
	Balance float64 `json:"current"`
	Summary float64 `json:"withdrawn"`
}

type Account struct {
	ID      uint `gorm:"primaryKey" sql:"AUTO_INCREMENT"`
	UserID  uint
	Balance sql.NullFloat64
}
