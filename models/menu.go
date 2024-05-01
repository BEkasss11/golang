package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	ID          int     `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"type:text"`
	Quantity    uint    `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	LastUpdated string  `gorm:"default:CURRENT_TIMESTAMP"`
}
