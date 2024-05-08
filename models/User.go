package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Username string `gorm:"unique;not null"`
	Role     string `gorm:"not null"`
}

type SignUpInput struct {
	Email    string `gorm:"unique;not null"`
	Password string `json:"password"`
	UserName string `json:"firstname"`
}

type SignInInput struct {
	Email    string `gorm:"unique;not null"`
	Password string `json:"password"`
	UserName string `json:"firstname"`
}
