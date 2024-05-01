package models

import (
	"github.com/BEkasss11/golang/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
    ID           uint   `gorm:"primaryKey"`
    Username     string `gorm:"unique;not null"`
    PasswordHash string `gorm:"not null"`
    Email        string `gorm:"unique;not null"`
    Role         string `gorm:"not null"`
}

type SignUpInput struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phoneNumber" gorm:"unique"`
	Password    string `json:"password"`

	Gender enums.Gender `json:"gender"`
}

type SignInInput struct {
	PhoneNumber string `json:"phoneNumber" gorm:"unique"`
	Password    string `json:"password"`
}
