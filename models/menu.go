package models

type Menu struct {
    ID          uint    `gorm:"primaryKey"`
    Name        string  `gorm:"not null"`
    Description string  `gorm:"type:text"`
	Quantity     uint    `gorm:"not null"`
    Price       float64 `gorm:"not null"`
	LastUpdated  string  `gorm:"default:CURRENT_TIMESTAMP"`
}

type Payment struct {
    ID               uint    `gorm:"primaryKey"`
    OrderID          uint    `gorm:"not null"`
    Amount           float64 `gorm:"not null"`
    PaymentTime      string  `gorm:"default:CURRENT_TIMESTAMP"`
    CardNumber       string  `gorm:"not null"`
    CardExpiryMonth  int     `gorm:"not null"`
    CardExpiryYear   int     `gorm:"not null"`
    CardCVV          string  `gorm:"not null"`
}