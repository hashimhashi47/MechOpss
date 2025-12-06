package models

import (
	"time"

	"gorm.io/gorm"
)



type Slot struct {
	gorm.Model
	CarModel  string    `json:"car"`
	CarNumber string    `json:"carnumber"`
	Problem   string    `json:"problem"`
	Time      time.Time `json:"time"`
	StaffName string    `json:"staffname"`
	StaffID   *uint     `json:"staff_id"`
	Status    string    `json:"status"`
}
