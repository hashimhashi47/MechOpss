package models

import "gorm.io/gorm"

// user
type User struct {
	gorm.Model
	FirstName    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email" gorm:"unique"`
	Role         string    `json:"role"`
	Phone        string    `json:"phone"`
	Bookings     []Booking `gorm:"foreignKey:UserID"`
	Booked       []Bookeds `gorm:"foreignKey:UserID"`
	Block        bool      `json:"block" gorm:"default:false"`
	Password     string    `json:"password" gorm:"type:varchar(255)"`
	RefreshToken string    `json:"-" gorm:"type:text"`
}
