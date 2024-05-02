package models

import "time"

type Order struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	OrderTime   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Status      string  `gorm:"not null"`
	TotalAmount float64 `gorm:"not null"`
}


type OrderItem struct {
	ID       uint   `gorm:"primaryKey"`
	menu_items  uint   `gorm:"not null"`
	ItemName string `gorm:"not null"`
	Quantity uint   `gorm:"not null"`
	Price    float64 `gorm:"not null"`
}
