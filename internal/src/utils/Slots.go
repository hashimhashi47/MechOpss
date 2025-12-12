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
			{Model: gorm.Model{ID: 1}, Status: "empty"},
			{Model: gorm.Model{ID: 2}, Status: "empty"},
			{Model: gorm.Model{ID: 3}, Status: "empty"},
			{Model: gorm.Model{ID: 4}, Status: "empty"},
			{Model: gorm.Model{ID: 5}, Status: "empty"},
		}

		for _, slot := range slots {
			if err := DB.Create(&slot).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
