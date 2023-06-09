package domain

type Order struct {
	ID         uint    `gorm:"primaryKey" sql:"AUTO_INCREMENT" json:"-"`
	UserID     uint    `json:"-"`
	Number     string  `gorm:"uniqueIndex:idx_numbers,sort:desc" json:"number"`
	Status     string  `json:"status"`
	Accrual    float64 `json:"accrual,omitempty"`
	UploadedAt int64   `gorm:"autoCreateTime" json:"uploaded_at"`
}

type OrderResponse struct {
	Number     string  `json:"number"`
	Status     string  `json:"status"`
	Accrual    float64 `json:"accrual,omitempty"`
	UploadedAt string  `json:"uploaded_at"`
}

type OrderLog struct {
	ID          uint    `gorm:"primaryKey" sql:"AUTO_INCREMENT" json:"-"`
	UserID      uint    `json:"-"`
	OrderNumber string  `json:"order"`
	Sum         float64 `json:"sum"`
	ProcessedAt int64   `gorm:"autoCreateTime" json:"processed_at"`
}

type OrderLogResponse struct {
	OrderNumber string  `json:"order"`
	Sum         float64 `json:"sum"`
	ProcessedAt string  `json:"processed_at"`
}
