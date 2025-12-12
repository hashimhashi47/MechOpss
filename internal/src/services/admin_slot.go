package services

import (
	"MechOpss/internal/src/models"
	"errors"
)


//get all the slots 
func (s *AdminService) ServiceGetAllSlots() ([]models.Slot, error) {
	var Slots []models.Slot
	if err := s.Repo.FindAll(&Slots); err != nil {
		return nil, errors.New("")
	}
	return Slots, nil
}


//count of total
func (s *AdminService) ServiceCountOfSlot() (int64, error) {
	Count, err := s.Repo.Count(&models.Slot{})
	if err != nil {
		return 0, errors.New("failed to get the slots count")
	}
	return Count, nil
}

//update the slot details
func (s *AdminService) ServiceUpdateSlot(input models.Slot, id string) (interface{}, error) {

	var Slot models.Slot
	if err := s.Repo.FindByID(&Slot, id); err != nil {
		return nil, errors.New("failed to find the slot")
	}

	if input.ServiceStart != "" {
		Slot.ServiceStart = input.ServiceStart
	}

	if input.ServiceEnd != "" {
		Slot.ServiceEnd = input.ServiceEnd
	}

	if input.Description != "" {
		Slot.Description = input.Description
	}

	if input.Status != "" {
		Slot.Status = input.Status
	}

	if err := s.Repo.Save(&Slot); err != nil {
		return nil, errors.New("unable to update the slot")
	}

	return Slot, nil
}

func (s *AdminService) ServiceRemoveSlot(id string) error {
	if err := s.Repo.Delete(models.Slot{}, id); err != nil {
		return err
	}
	return nil
}

//empty the slot after the completion
func (s *AdminService) ServiceEmptySlot(id string) (interface{}, error) {
	var slot models.Slot
	if err := s.Repo.FindByID(&slot, id); err != nil {
		return nil, errors.New("failed to find the slot")
	}

	slot.CarModel = ""
	slot.Time = nil
	slot.CarNumber = ""
	slot.ServiceStart = ""
	slot.ServiceEnd = ""
	slot.Description = ""
	slot.StaffID = nil
	slot.StaffName = ""
	slot.Status = "empty"

	if err := s.Repo.Save(&slot); err != nil {
		return nil, errors.New("unable to remove the slot")
	}

	return slot, nil
}
