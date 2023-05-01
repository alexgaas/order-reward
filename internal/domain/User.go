package domain

type User struct {
	ID       uint   `gorm:"primaryKey" sql:"AUTO_INCREMENT" json:"-"`
	Login    string `gorm:"uniqueIndex:idx_logins" json:"login"`
	Password string `json:"password,omitempty"`
}
