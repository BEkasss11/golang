package models

type Menu struct {
    ID          uint    `gorm:"primaryKey"`
    Name        string  `gorm:"not null"`
    Description string  `gorm:"type:text"`
	Quantity     uint    `gorm:"not null"`
    Price       float64 `gorm:"not null"`
	LastUpdated  string  `gorm:"default:CURRENT_TIMESTAMP"`
}

