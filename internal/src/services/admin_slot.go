package services

import (
	"MechOpss/internal/src/models"
	"errors"
)

func (s *AdminService) ServiceGetAllSlots() ([]models.Slot, error) {
	var Slots []models.Slot

	if err := s.Repo.FindAll(&Slots); err != nil {
		return nil, errors.New("")
	}
	return Slots, nil
}


