package models

import "gorm.io/gorm"


type Admin struct {
	gorm.Model
	Name         string
	Email        string `json:"email" gorm:"unique"`
	Role         string `gorm:"default:admin"`
	RefreshToken string `json:"-" gorm:"type:text"`
	Password     string `json:"password"`
}


