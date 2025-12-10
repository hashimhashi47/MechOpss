package models

import (
	"time"

	"gorm.io/gorm"
)

type Slot struct {
	gorm.Model
	CarModel     string    `json:"car"`
	Time         *time.Time `json:"time"`
	CarNumber    string    `json:"carnumber"`
	ServiceStart string    `json:"service_start"`
	ServiceEnd   string    `json:"service_end"`
	Description  string    `json:"description"`

	StaffID   *uint     `json:"staff_id"`
	StaffName string    `json:"staff_name"`
	Status    string    `json:"status"`
}
