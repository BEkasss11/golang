package models

type Payment struct {
	ID              uint    `gorm:"primaryKey"`
	OrderID         uint    `gorm:"not null"`
	Amount          float64 `gorm:"not null"`
	PaymentTime     string  `gorm:"default:CURRENT_TIMESTAMP"`
	CardNumber      string  `gorm:"not null"`
	CardExpiryMonth int     `gorm:"not null"`
	CardExpiryYear  int     `gorm:"not null"`
	CardCVV         string  `gorm:"not null"`
}
