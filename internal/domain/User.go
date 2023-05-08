package domain

type User struct {
	ID       uint    `gorm:"primaryKey" sql:"AUTO_INCREMENT" json:"-"`
	Login    string  `gorm:"uniqueIndex:idx_logins" json:"login"`
	Password string  `json:"password,omitempty"`
	Account  Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`

	Orders    []Order    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	OrderList []OrderLog `json:"-"`
}

type LoginResponse struct {
	Authtoken string `json:"auth_token"`
}
