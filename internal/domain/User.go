package domain

type User struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	Login    string `gorm:"uniqueIndex:idx_logins" json:"login"`
	Password string `json:"password,omitempty"`
}
