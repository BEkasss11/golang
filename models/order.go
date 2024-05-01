package models

type Order struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	OrderTime   string  `gorm:"default:CURRENT_TIMESTAMP"`
	Status      string  `gorm:"not null"`
	TotalAmount float64 `gorm:"not null"`
}

type OrderItem struct {
	ID       uint   `gorm:"primaryKey"`
	OrderID  uint   `gorm:"not null"`
	ItemName string `gorm:"not null"`
	Quantity uint   `gorm:"not null"`
	Price    float64 `gorm:"not null"`
	NotType  string
}
