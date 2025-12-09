package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	FirstName    string   `json:"firstname"`
	LastName     string   `json:"lastname"`
	Email        string   `json:"email" gorm:"unique"`
	Role         string   `json:"role"`
	IdentityCard string   `json:"cardnumber"`
	Department   string   `json:"department"`
	Block        bool     `json:"block" gorm:"default:false"`
	RefreshToken string   `json:"-" gorm:"type:text"`
	Password     string   `json:"password"`
	Bookings     []Bookeds `gorm:"foreignKey:StaffID;references:ID"`
}
