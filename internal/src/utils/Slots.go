package utils

import (
	"MechOpss/internal/src/models"

	"gorm.io/gorm"
)



func SeedSlots(DB *gorm.DB) error {
	var count int64
	DB.Model(&models.Slot{}).Count(&count)

	if count == 0 {
		slots := []models.Slot{
			{Model: gorm.Model{ID: 1}, Status: "EMPTY"},
			{Model: gorm.Model{ID: 2}, Status: "EMPTY"},
			{Model: gorm.Model{ID: 3}, Status: "EMPTY"},
			{Model: gorm.Model{ID: 4}, Status: "EMPTY"},
			{Model: gorm.Model{ID: 5}, Status: "EMPTY"},
		}

		for _, slot := range slots {
			if err := DB.Create(&slot).Error; err != nil {
				return err
			}
		}
	}

	return nil
}