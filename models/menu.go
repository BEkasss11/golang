package models

type Menu struct {
	ID          int     `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"type:text"`
    Price       float64 `gorm:"not null"`
	Quantity    uint    `gorm:"not null"`
	LastUpdated string  `gorm:"т"`
}
